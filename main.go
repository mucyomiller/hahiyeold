package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	jwt "github.com/dgrijalva/jwt-go"
	pb "github.com/mucyomiller/hahiye/hahiye"
	"github.com/mucyomiller/hahiye/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	port        = ":9090"
	srvCertFile = "./certs/server.crt"
	srvKeyFile  = "./certs/ca-key.pem"
)

func main() {

	//spin up dgraph db connection
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
		return
	}
	fmt.Println("Now Connected to DgraphDB")
	defer conn.Close()
	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)
	op := &api.Operation{}
	// Adding Our Defined  Database SCHEMA
	op.Schema = `
	location: geo @index(geo) .
	name: string @index(term) .
	lastname: string @index(term) .
	email: string @index(exact) @upsert .
	password: password .
	follows: uid @reverse .
	interested: uid @reverse .
	verified: bool .
	`
	ctx := context.Background()
	err = dg.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

	// creating a listener on specified port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}

	fmt.Println("starting gprc server...")

	// creating tls credentials from cert file
	tlsCreds, err := credentials.NewServerTLSFromFile(srvCertFile, srvKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// setup and register our available services
	authService := server.NewAuthService(dg)
	accountService := server.NewAccountServiceServer(dg)
	placeService := server.NewPlaceServiceServer(dg)
	interestService := server.NewInterestServiceServer(dg)

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCreds),
		grpc.UnaryInterceptor(authUnaryIntercept),
		grpc.StreamInterceptor(streamAuthIntercept),
	)

	pb.RegisterAuthServiceServer(grpcServer, authService)
	pb.RegisterAccountServiceServer(grpcServer, accountService)
	pb.RegisterPlaceServiceServer(grpcServer, placeService)
	pb.RegisterInterestServiceServer(grpcServer, interestService)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			grpcServer.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start services server
	log.Println("starting secure rpc services on", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

}

// authUnaryIntercept intercepts incoming requests to validate
// jwt token from metadata header "authorization"
func authUnaryIntercept(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	//bypass auth if method is /hahiye.AuthService/Login
	if info.FullMethod == "/hahiye.AuthService/Login" {
		fmt.Println("bypassing auth cz it's login action")
		return handler(ctx, req)
	}
	if err := auth(ctx); err != nil {
		return nil, err
	}
	log.Println("authorization OK")
	return handler(ctx, req)
}

// streamAuthIntercept intercepts to validate authorization
func streamAuthIntercept(
	server interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	//bypass auth if method is /hahiye.AuthService/Login
	if info.FullMethod == "/hahiye.AuthService/Login" {
		fmt.Println("bypassing auth cz it's login action")
		return handler(server, stream)
	}
	if err := auth(stream.Context()); err != nil {
		return err
	}
	log.Println("authorization OK")
	return handler(server, stream)
}

func auth(ctx context.Context) error {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(
			codes.InvalidArgument,
			"missing context",
		)
	}

	authString, ok := meta["authorization"]
	if !ok {
		return status.Errorf(
			codes.Unauthenticated,
			"missing authorization",
		)
	}
	// validate token algo
	log.Println("found jwt token")
	jwtToken, err := jwt.Parse(
		authString[0],
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("bad signing method")
			}
			// additional validation goes here.
			return []byte("s3cr3t"), nil
		},
	)

	if jwtToken.Valid {
		return nil
	}
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return status.Error(codes.Internal, "bad token")
}
