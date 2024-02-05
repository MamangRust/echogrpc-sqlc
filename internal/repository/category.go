package repository

import (
	"MamangRust/echobloggrpc/internal/domain/requests"
	db "MamangRust/echobloggrpc/pkg/database/postgres/schema"
	"context"
	"errors"
)

type categoryRepository struct {
	db      *db.Queries
	context context.Context
}

func NewCategoryRepository(db *db.Queries, context context.Context) *categoryRepository {
	return &categoryRepository{db: db, context: context}
}

func (r *categoryRepository) FindAll() ([]*db.Category, error) {
	res, err := r.db.GetCategories(r.context)

	if err != nil {
		return nil, errors.New("failed get categories")
	}

	return res, nil
}

func (r *categoryRepository) FindByID(id int) (*db.Category, error) {

	res, err := r.db.GetCategory(r.context, int32(id))

	if err != nil {
		return nil, errors.New("failed get category by id")
	}

	return res, nil
}

func (r *categoryRepository) Create(input *requests.CreateCategoryRequest) (*db.Category, error) {

	res, err := r.db.CreateCategory(r.context, input.Name)

	if err != nil {
		return nil, errors.New("failed create category")
	}

	return res, nil
}

func (r *categoryRepository) Update(input *requests.UpdateCategoryRequest) (*db.Category, error) {
	var categoryRequest db.UpdateCategoryParams
	res, err := r.db.GetCategory(r.context, int32(input.ID))

	if err != nil {
		return nil, errors.New("failed get category by id")
	}

	categoryRequest.ID = res.ID
	categoryRequest.Name = input.Name

	updated, err := r.db.UpdateCategory(r.context, categoryRequest)

	if err != nil {
		return nil, errors.New("failed update category")
	}

	return updated, nil
}

func (r *categoryRepository) Delete(id int) error {
	_, err := r.db.GetCategory(r.context, int32(id))

	if err != nil {
		return errors.New("failed get category by id")
	}

	err = r.db.DeleteCategory(r.context, int32(id))

	if err != nil {
		return errors.New("failed delete category")
	}

	return nil
}
