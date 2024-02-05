package service

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	"MamangRust/echobloggrpc/internal/repository"
	db "MamangRust/echobloggrpc/pkg/database/postgres/schema"
	"MamangRust/echobloggrpc/pkg/logger"

	"go.uber.org/zap"
)

type commentService struct {
	logger     logger.Logger
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository, logger logger.Logger) *commentService {
	return &commentService{repository: repository, logger: logger}
}

func (s *commentService) FindAll() ([]*db.Comment, error) {
	res, err := s.repository.FindAll()

	if err != nil {
		s.logger.Error("Error fetching comments: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *commentService) FindById(id int) (*db.Comment, error) {
	res, err := s.repository.FindByID(id)

	if err != nil {
		s.logger.Error("Error fetching comment: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *commentService) Create(input *requests.CreateCommentRequest) (*db.Comment, error) {
	res, err := s.repository.Create(input)

	if err != nil {
		s.logger.Error("Error creating comment: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *commentService) Update(input *requests.UpdateCommentRequest) (*db.Comment, error) {
	res, err := s.repository.Update(input)

	if err != nil {
		s.logger.Error("Error updating comment: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *commentService) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		s.logger.Error("Error delete comment: ", zap.Error(err))
	}

	return err
}
