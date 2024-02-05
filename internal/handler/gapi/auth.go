package gapi

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/pb"
	"MamangRust/echobloggrpc/internal/service"
	"context"
)

type authHandleGrpc struct {
	pb.UnimplementedAuthServiceServer
	auth service.AuthService
}

func NewAuthHandlerGrpc(auth service.AuthService) *authHandleGrpc {
	return &authHandleGrpc{auth: auth}
}

func (a *authHandleGrpc) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := requests.CreateUserRequest{
		FirstName:       req.Firstname,
		LastName:        req.Lastname,
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}

	res, err := a.auth.Register(&user)

	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		User: &pb.User{
			Firstname: res.Firstname,
			Lastname:  res.Lastname,
			Email:     res.Email,
		},
	}, nil
}

func (a *authHandleGrpc) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := requests.AuthLoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := a.auth.Login(&user)

	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Status: "Success",
		Token:  res.Token,
	}, nil
}
