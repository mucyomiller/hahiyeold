package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mucyomiller/hahiye/hahiye"
)

// AccountService implements the pb AccountServiceServer  interface
type AccountService struct{}

func newAccountServiceServer() pb.AccountServiceServer {
	return new(AccountService)
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
