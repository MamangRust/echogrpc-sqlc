package service

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/repository"
	"MamangRust/echobloggrpc/pkg/auth"
	db "MamangRust/echobloggrpc/pkg/database/postgres/schema"
	"MamangRust/echobloggrpc/pkg/hash"

	"errors"
)

type authService struct {
	auth  repository.UserRepository
	hash  hash.Hashing
	token auth.TokenManager
}

func NewAuthService(auth repository.UserRepository, hash hash.Hashing, token auth.TokenManager) *authService {
	return &authService{auth: auth, hash: hash, token: token}
}

func (s *authService) Register(request *requests.CreateUserRequest) (*db.User, error) {
	_, err := s.auth.FindByEmail(request.Email)

	if err == nil {
		return nil, errors.New("failed email already exist")
	}

	passwordHash, err := s.hash.HashPassword(request.Password)

	if err != nil {
		return nil, err
	}

	user := requests.CreateUserRequest{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  passwordHash,
	}

	res, err := s.auth.Create(&user)

	if err != nil {
		return nil, errors.New("failed create user :" + err.Error())
	}

	return res, nil

}

func (s *authService) Login(request *requests.AuthLoginRequest) (*requests.JWTToken, error) {
	res, err := s.auth.FindByEmail(request.Email)

	if err != nil {
		return nil, errors.New("failed get user " + err.Error())
	}

	err = s.hash.ComparePassword(res.Password, request.Password)

	if err != nil {
		return nil, errors.New("failed compare password " + err.Error())
	}

	token, err := s.createJwt(res.Firstname+" "+res.Lastname, res.ID)

	if err != nil {
		return nil, err
	}

	return &requests.JWTToken{
		Token: token,
	}, nil
}

func (s *authService) createJwt(fullname string, id int32) (string, error) {
	token, err := s.token.GenerateToken(fullname, id)

	if err != nil {
		return "", err
	}

	return token, nil
}
