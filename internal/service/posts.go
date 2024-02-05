package service

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/repository"
	db "MamangRust/echobloggrpc/pkg/database/postgres/schema"
	"MamangRust/echobloggrpc/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

type postService struct {
	logger     logger.Logger
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository, logger logger.Logger) *postService {
	return &postService{repository: repository, logger: logger}
}

func (s *postService) FindAll() ([]*db.Post, error) {

	res, err := s.repository.FindAll()

	if err != nil {
		s.logger.Error("Error fetching posts: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *postService) FindById(id int) (*db.GetPostRow, error) {
	res, err := s.repository.FindByID(id)

	if err != nil {
		s.logger.Error("Error fetching post: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *postService) FindByIDRelationJoin(id int) (*db.GetPostRelationRow, error) {
	res, err := s.repository.FindByIDRelationJoin(id)

	if err != nil {
		s.logger.Error("Error fetching post: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *postService) Create(request *requests.CreatePostRequest) (*db.Post, error) {
	var post requests.CreatePostRequest

	post.Title = request.Title
	post.Slug = request.Slug
	post.Img = request.Img
	post.Body = request.Body
	post.CategoryID = request.CategoryID
	post.UserID = request.UserID
	post.UserName = request.UserName

	res, err := s.repository.Create(&post)

	if err != nil {
		s.logger.Error("Error creating post: ", zap.Error(err))
		return nil, err
	}

	return res, err
}

func (s *postService) Update(request *requests.UpdatePostRequest) (*db.Post, error) {
	var post requests.UpdatePostRequest

	post.ID = request.ID
	post.Title = request.Title
	post.Slug = request.Slug
	post.Img = request.Img
	post.Body = request.Body
	post.CategoryID = request.CategoryID
	post.UserID = request.UserID
	post.UserName = request.UserName

	fmt.Println(post)

	res, err := s.repository.Update(&post)

	if err != nil {
		s.logger.Error("Error updating post: ", zap.Error(err))
		return nil, err
	}

	return res, err
}

func (s *postService) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		s.logger.Error("Error delete post: ", zap.Error(err))

		return err
	}

	return nil

}
