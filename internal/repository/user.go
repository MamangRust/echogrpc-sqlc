package repository

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	db "MamangRust/echobloggrpc/pkg/database/postgres/schema"
	"context"
	"errors"
	"fmt"
)

type userRepository struct {
	db      *db.Queries
	context context.Context
}

func NewUserRepository(db *db.Queries, context context.Context) *userRepository {
	return &userRepository{db: db, context: context}
}

func (r *userRepository) FindAll() ([]*db.GetUsersRow, error) {

	user, err := r.db.GetUsers(r.context)

	if err != nil {
		return nil, errors.New("failed get users")
	}

	return user, nil
}

func (r *userRepository) FindById(id int) (*db.GetUserRow, error) {
	user, err := r.db.GetUser(r.context, int32(id))

	if err != nil {
		return nil, errors.New("failed get user")
	}

	return user, nil
}

func (r *userRepository) Create(input *requests.CreateUserRequest) (*db.User, error) {
	var userRequest db.CreateUserParams

	userRequest.Firstname = input.FirstName
	userRequest.Lastname = input.LastName
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	user, err := r.db.CreateUser(r.context, userRequest)

	if err != nil {
		return nil, errors.New("failed create user")
	}

	return user, nil
}

func (r *userRepository) Update(input *requests.UpdateUserRequest) (*db.User, error) {
	var userRequest db.UpdateUserParams

	_, err := r.db.GetUser(r.context, int32(input.ID))

	if err != nil {
		return nil, errors.New("failed get user")
	}

	userRequest.Firstname = input.FirstName
	userRequest.Lastname = input.LastName
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	res, err := r.db.UpdateUser(r.context, userRequest)

	if err != nil {
		return nil, errors.New("failed update user")
	}

	return res, nil

}

func (r *userRepository) Delete(id int) error {
	resid, err := r.db.GetUser(r.context, int32(id))

	if err != nil {
		return errors.New("failed get user")
	}

	err = r.db.DeleteUser(r.context, resid.ID)

	if err != nil {
		return fmt.Errorf("failed error")
	}

	return nil

}

func (r *userRepository) FindByEmail(email string) (*db.FindByEmailUserRow, error) {
	res, err := r.db.FindByEmailUser(r.context, email)

	if err != nil {
		return nil, errors.New("failed get user")
	}

	return res, nil
}
