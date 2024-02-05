package requests

import "github.com/go-playground/validator/v10"

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func (r *CreateCategoryRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil

}

type UpdateCategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (r *UpdateCategoryRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil

}
