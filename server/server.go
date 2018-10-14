package server

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/dgraph-io/dgo"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mucyomiller/hahiye/hahiye"
	"github.com/mucyomiller/hahiye/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AccountService implements the pb AccountServiceServer  interface
type AccountService struct {
	db *dgo.Dgraph
}

// NewAccountServiceServer create instance of AccountService
func NewAccountServiceServer(db *dgo.Dgraph) pb.AccountServiceServer {
	return &AccountService{db: db}
}

// CreateAccount used to create new user account
func (a *AccountService) CreateAccount(context.Context, *pb.Account) (*pb.Account, error) {
	return &pb.Account{}, nil
}

// DeleteAccount used to delete user account
func (a *AccountService) DeleteAccount(context.Context, *pb.AccountRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

// GetAccount used to get single account
func (a *AccountService) GetAccount(context.Context, *pb.AccountRequest) (*pb.Account, error) {

	return &pb.Account{}, nil
}

// UpdateAccount used to update user account info
func (a *AccountService) UpdateAccount(context.Context, *pb.Account) (*pb.AccountResponse, error) {
	return &pb.AccountResponse{}, nil
}

// PlaceService implements the pb PlaceServiceServer  interface
type PlaceService struct {
	db *dgo.Dgraph
}

// NewPlaceServiceServer create instance of PlaceService
func NewPlaceServiceServer(db *dgo.Dgraph) pb.PlaceServiceServer {
	return &PlaceService{db: db}
}

// AddPlace adds new place
func (p *PlaceService) AddPlace(context.Context, *pb.Place) (*pb.Place, error) {
	return &pb.Place{}, nil
}

// DeletePlace delete specified place
func (p *PlaceService) DeletePlace(context.Context, *pb.PlaceRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

// GetPlace return info about specified place
func (p *PlaceService) GetPlace(context.Context, *pb.PlaceRequest) (*pb.Place, error) {
	return &pb.Place{}, nil
}

// GetPlaces stream available places
func (p *PlaceService) GetPlaces(*empty.Empty, pb.PlaceService_GetPlacesServer) error {
	return nil
}

// UpdatePlace update info about specified place
func (p *PlaceService) UpdatePlace(context.Context, *pb.Place) (*pb.Place, error) {
	return &pb.Place{}, nil
}

// InterestService implements the pb InterestServiceServer  interface
type InterestService struct {
	db *dgo.Dgraph
}

// NewInterestServiceServer create instance of InterestService
func NewInterestServiceServer(db *dgo.Dgraph) pb.InterestServiceServer {
	return &InterestService{db: db}
}

// AddInterest adds new Interest
func (i *InterestService) AddInterest(context.Context, *pb.Interest) (*pb.Interest, error) {
	return &pb.Interest{}, nil
}

// Removeinterest specified Interest
func (i *InterestService) Removeinterest(context.Context, *pb.InterestRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

// GetInterest specified Interest
func (i *InterestService) GetInterest(ctx context.Context, req *pb.InterestRequest) (*pb.Interest, error) {
	resp, err := i.db.NewTxn().QueryWithVars(ctx,
		`query interest($id: string = "0", $name: string = "") {
			interest(func: has(interest)) @filter(uid($id) OR allofterms(name, $name)) {
				interest
				uid
				name
			}
		  }`, map[string]string{"$id": req.GetId(), "$name": req.GetName()})

	if err != nil {
		//TODO: return grpc error
		return nil, status.Errorf(
			codes.InvalidArgument, err.Error(),
		)
	}

	// InterestResp to hold Interest response
	type InterestResp struct {
		Interest []model.Interest
	}
	data := InterestResp{}
	// Unmarshal this form {"interest":[{"interest":"","uid":"0x1","name":"chips"}]}
	err = json.Unmarshal(resp.Json, &data)
	if err != nil {
		return nil, status.Errorf(
			codes.Unknown, err.Error(),
		)
	}
	return &pb.Interest{Id: data.Interest[0].UID, Name: data.Interest[0].Name}, nil
}

// GetInterests stream available Interests
func (i *InterestService) GetInterests(empty *empty.Empty, stream pb.InterestService_GetInterestsServer) error {
	resp, err := i.db.NewTxn().Query(context.Background(),
		`{
			interests(func: has(interest)){
			  uid
			  name
			}
		}`)
	if err != nil {
		return status.Errorf(codes.Unknown, err.Error())
	}

	type InterestsResp struct {
		Interests []model.Interest `json:"interests"`
	}
	data := InterestsResp{}
	// unmarshall {"interests": [{"uid": "0x1","name": "chips"},{"uid": "0x3","name": "chips"}]}
	err = json.Unmarshal(resp.Json, &data)
	if err != nil {
		return status.Errorf(codes.Unknown, err.Error())
	}
	// try to  stream all found interests
	for _, interest := range data.Interests {
		s := &pb.Interest{
			Id:   interest.UID,
			Name: interest.Name,
		}
		err := stream.Send(s)
		if err != nil {
			return status.Errorf(codes.Unknown, err.Error())
		}
	}
	return nil
}

// UpdateInterest update specified Interest
func (i *InterestService) UpdateInterest(context.Context, *pb.Interest) (*pb.InterestResponse, error) {
	return &pb.InterestResponse{}, nil
}

type user struct {
	username string
	password []byte
}

// AuthService implement AuthServiceServer
type AuthService struct {
	db *dgo.Dgraph
}

// NewAuthService returns *AuthService
func NewAuthService(db *dgo.Dgraph) *AuthService {
	return &AuthService{db: db}
}

// Login credentials are stored in dgraph
// username and hashed password.
func (s *AuthService) Login(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Println("Authorizing user", req.GetUsername())
	if req.GetUsername() == "" || req.GetPassword() == "" {
		log.Println("Auth failed for user", req.GetUsername())
		return nil, status.Errorf(codes.InvalidArgument, "missing username or password")
	}
	// query dgraph here
	// temporary mocks user credentials
	secret := "s3cr3t"
	name := "Mucyo Miller"
	username := "miller"
	password := "miller"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if req.GetUsername() != username {
		log.Println("missing username")
		return nil, status.Error(codes.PermissionDenied, "invalid user")
	}

	if err := bcrypt.CompareHashAndPassword(hash, []byte(req.GetPassword())); err != nil {
		log.Println("auth failed")
		return nil, status.Error(codes.PermissionDenied, "auth failed")
	}

	// create jwt token
	// see reserved claims https://tools.ietf.org/html/rfc7519#section-4.1
	// see jwt example here https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":  time.Now().Add(time.Minute * 20).Unix(),
			"sub":  username,
			"iss":  "authservice",
			"aud":  "user",
			"name": name,
		},
	)

	// this example uses a simple string secret. You can also
	// use JWT package to specify an RSA public cert here as well.
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal login problem")
	}

	log.Printf("User %s logged in OK, JWT token: %s\n", username, tokenString)
	return &pb.AuthResponse{Token: tokenString}, nil
}
