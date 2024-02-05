package service

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	db "MamangRust/echobloggrpc/pkg/database/postgres/schema"
)

type CategoryService interface {
	FindAll() ([]*db.Category, error)
	FindByID(id int) (*db.Category, error)
	Create(input *requests.CreateCategoryRequest) (*db.Category, error)
	Update(input *requests.UpdateCategoryRequest) (*db.Category, error)
	Delete(id int) error
}

type PostService interface {
	FindAll() ([]*db.Post, error)
	FindById(id int) (*db.GetPostRow, error)
	FindByIDRelationJoin(id int) (*db.GetPostRelationRow, error)
	Create(requests *requests.CreatePostRequest) (*db.Post, error)
	Update(requests *requests.UpdatePostRequest) (*db.Post, error)
	Delete(id int) error
}

type CommentService interface {
	FindAll() ([]*db.Comment, error)
	FindById(id int) (*db.Comment, error)
	Create(input *requests.CreateCommentRequest) (*db.Comment, error)
	Update(input *requests.UpdateCommentRequest) (*db.Comment, error)
	Delete(id int) error
}

type UserService interface {
	FindAll() ([]*db.GetUsersRow, error)
	FindByID(id int) (*db.GetUserRow, error)
	Create(requests *requests.CreateUserRequest) (*db.User, error)
	Update(requests *requests.UpdateUserRequest) (*db.User, error)
	DeleteId(id int) error
}

type AuthService interface {
	Register(requests *requests.CreateUserRequest) (*db.User, error)
	Login(request *requests.AuthLoginRequest) (*requests.JWTToken, error)
}
