package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mucyomiller/hahiye/hahiye"
)

// AccountService implements the pb AccountServiceServer  interface
type AccountService struct{}

// NewAccountServiceServer create instance of AccountService
func NewAccountServiceServer() pb.AccountServiceServer {
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

// PlaceService implements the pb PlaceServiceServer  interface
type PlaceService struct{}

// NewPlaceServiceServer create instance of PlaceService
func NewPlaceServiceServer() pb.PlaceServiceServer {
	return new(PlaceService)
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
type InterestService struct{}

// NewInterestServiceServer create instance of InterestService
func NewInterestServiceServer() pb.InterestServiceServer {
	return new(InterestService)
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
func (i *InterestService) GetInterest(context.Context, *pb.InterestRequest) (*pb.Interest, error) {
	return &pb.Interest{}, nil
}

// GetInterests stream available Interests
func (i *InterestService) GetInterests(*empty.Empty, pb.InterestService_GetInterestsServer) error {
	return nil
}

// UpdateInterest update specified Interest
func (i *InterestService) UpdateInterest(context.Context, *pb.Interest) (*pb.InterestResponse, error) {
	return &pb.InterestResponse{}, nil
}
