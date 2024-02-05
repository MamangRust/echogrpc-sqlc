package service

import (
	"MamangRust/echobloggrpc/internal/repository"
	"MamangRust/echobloggrpc/pkg/auth"
	"MamangRust/echobloggrpc/pkg/hash"
	"MamangRust/echobloggrpc/pkg/logger"
	"errors"
)

type Services struct {
	Auth     AuthService
	User     UserService
	Category CategoryService
	Post     PostService
	Comment  CommentService
}

type Deps struct {
	Repository *repository.Repositories
	Logger     *logger.Logger
	Hash       *hash.Hashing
	Token      auth.TokenManager
}

func NewServices(deps Deps) (*Services, error) {
	if deps.Repository.User == nil || deps.Hash == nil || deps.Logger == nil || deps.Token == nil {
		return nil, errors.New("nil dependency detected")
	}

	return &Services{
		Auth:     NewAuthService(deps.Repository.User, *deps.Hash, deps.Token),
		User:     NewUserService(deps.Repository.User, *deps.Hash, *deps.Logger),
		Category: NewCategoryService(deps.Repository.Category, *deps.Logger),
		Post:     NewPostService(deps.Repository.Post, *deps.Logger),
		Comment:  NewCommentService(deps.Repository.Comment, *deps.Logger),
	}, nil
}
