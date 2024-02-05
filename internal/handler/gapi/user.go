package gapi

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/pb"
	"MamangRust/echobloggrpc/internal/service"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userHandleGrpc struct {
	pb.UnimplementedUserServiceServer
	user service.UserService
}

func NewUserHandleGrpc(user service.UserService) *userHandleGrpc {
	return &userHandleGrpc{user: user}
}

func (h *userHandleGrpc) GetUsers(ctx context.Context, request *emptypb.Empty) (*pb.UsersResponse, error) {
	res, err := h.user.FindAll()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbUsers []*pb.User

	for _, user := range res {
		pbUsers = append(pbUsers, &pb.User{
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
		})
	}

	return &pb.UsersResponse{Users: pbUsers}, nil
}

func (h *userHandleGrpc) GetUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	res, err := h.user.FindByID(int(request.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserResponse{
		User: &pb.User{
			Firstname: res.Firstname,
			Lastname:  res.Lastname,
			Email:     res.Email,
		},
	}, nil
}

func (s *userHandleGrpc) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.UserResponse, error) {
	res, err := s.user.Create(&requests.CreateUserRequest{
		FirstName: request.Firstname,
		LastName:  request.Lastname,
		Email:     request.Email,
		Password:  request.Password,
	})

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		User: &pb.User{
			Firstname: res.Firstname,
			Lastname:  res.Lastname,
			Email:     res.Email,
		},
	}, nil

}

func (h *userHandleGrpc) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	res, err := h.user.Update(&requests.UpdateUserRequest{
		ID:        int(request.Id),
		FirstName: request.Firstname,
		LastName:  request.Lastname,
		Email:     request.Email,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserResponse{
		User: &pb.User{
			Firstname: res.Firstname,
			Lastname:  res.Lastname,
			Email:     res.Email,
		},
	}, nil
}

func (h *userHandleGrpc) DeleteUser(ctx context.Context, request *pb.UserRequest) (*pb.DeleteUserResponse, error) {
	err := h.user.DeleteId(int(request.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}
