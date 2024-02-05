package service

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/repository"
	db "MamangRust/echobloggrpc/pkg/database/postgres/schema"
	"MamangRust/echobloggrpc/pkg/hash"
	"MamangRust/echobloggrpc/pkg/logger"

	"go.uber.org/zap"
)

type userService struct {
	repository repository.UserRepository
	hash       hash.Hashing
	logger     logger.Logger
}

func NewUserService(user repository.UserRepository, hash hash.Hashing, logger logger.Logger) *userService {
	return &userService{repository: user, hash: hash, logger: logger}
}

func (s *userService) FindAll() ([]*db.GetUsersRow, error) {
	res, err := s.repository.FindAll()

	if err != nil {
		s.logger.Error("failed get users: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *userService) FindByID(id int) (*db.GetUserRow, error) {
	res, err := s.repository.FindById(id)

	if err != nil {
		s.logger.Error("failed get user by id: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *userService) Create(request *requests.CreateUserRequest) (*db.User, error) {
	var createRequest requests.CreateUserRequest

	hashing, err := s.hash.HashPassword(request.Password)

	if err != nil {
		s.logger.Error("Error hashing password: ", zap.Error(err))
		return nil, err
	}

	_, err = s.repository.FindByEmail(request.Email)

	if err != nil {
		s.logger.Error("Error fetching user: ", zap.Error(err))

		return nil, err
	}

	createRequest.FirstName = request.FirstName
	createRequest.LastName = request.LastName
	createRequest.Email = request.Email
	createRequest.Password = hashing

	user, err := s.repository.Create(&createRequest)

	if err != nil {
		s.logger.Error("Error creating user: ", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(request *requests.UpdateUserRequest) (*db.User, error) {
	var userRequest requests.UpdateUserRequest

	_, err := s.repository.FindById(userRequest.ID)

	if err != nil {
		s.logger.Error("Errror update user: ", zap.Error(err))

		return nil, err
	}

	hashing, err := s.hash.HashPassword(request.Password)

	if err != nil {
		s.logger.Error("Error hashing password: ", zap.Error(err))

		return nil, err
	}

	userRequest.ID = request.ID

	userRequest.FirstName = request.FirstName
	userRequest.LastName = request.LastName
	userRequest.Email = request.Email
	userRequest.Password = hashing

	res, err := s.repository.Update(&userRequest)

	if err != nil {
		s.logger.Error("Error update user: ", zap.Error(err))

		return nil, err
	}

	return res, nil
}

func (s *userService) DeleteId(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		s.logger.Error("Error delete user: ", zap.Error(err))

		return err
	}

	return nil
}
