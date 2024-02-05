package repository

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	db "MamangRust/echobloggrpc/pkg/database/postgres/schema"
)

type CategoryRepository interface {
	FindAll() ([]*db.Category, error)
	FindByID(id int) (*db.Category, error)
	Create(input *requests.CreateCategoryRequest) (*db.Category, error)
	Update(input *requests.UpdateCategoryRequest) (*db.Category, error)
	Delete(id int) error
}

type CommentRepository interface {
	FindAll() ([]*db.Comment, error)
	FindByID(id int) (*db.Comment, error)
	Create(input *requests.CreateCommentRequest) (*db.Comment, error)
	Update(input *requests.UpdateCommentRequest) (*db.Comment, error)
	Delete(id int) error
}

type PostRepository interface {
	FindAll() ([]*db.Post, error)
	FindByID(id int) (*db.GetPostRow, error)
	FindByIDRelationJoin(id int) (*db.GetPostRelationRow, error)
	Create(input *requests.CreatePostRequest) (*db.Post, error)
	Update(input *requests.UpdatePostRequest) (*db.Post, error)
	Delete(id int) error
}

type UserRepository interface {
	FindAll() ([]*db.GetUsersRow, error)
	FindByEmail(email string) (*db.FindByEmailUserRow, error)
	FindById(id int) (*db.GetUserRow, error)
	Create(input *requests.CreateUserRequest) (*db.User, error)
	Update(input *requests.UpdateUserRequest) (*db.User, error)
	Delete(id int) error
}
