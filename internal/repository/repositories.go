package repository

import (
	db "MamangRust/echobloggrpc/pkg/database/postgres/schema"
	"context"
)

type Repositories struct {
	User     UserRepository
	Category CategoryRepository
	Post     PostRepository
	Comment  CommentRepository
}

func NewRepositories(db *db.Queries, context context.Context) *Repositories {
	return &Repositories{
		User:     NewUserRepository(db, context),
		Category: NewCategoryRepository(db, context),
		Post:     NewPostsRepository(db, context),
		Comment:  NewCommentRepository(db, context),
	}
}
