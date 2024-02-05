package requests

import "github.com/go-playground/validator/v10"

type CreatePostRequest struct {
	Title      string `json:"title" form:"title"`
	Slug       string `json:"slug" form:"slug"`
	Img        string `json:"img" form:"img"`
	Body       string `json:"body" form:"body"`
	CategoryID int    `json:"category_id" form:"category_id"`
	UserID     int    `json:"user_id" form:"user_id"`
	UserName   string `json:"user_name" form:"user_name"`
}

func (u *CreatePostRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	if err != nil {
		return err
	}

	return nil
}

type UpdatePostRequest struct {
	ID         int    `json:"id" form:"id"`
	Title      string `json:"title" form:"title"`
	Slug       string `json:"slug" form:"slug"`
	Img        string `json:"img" form:"img"`
	Body       string `json:"body" form:"body"`
	CategoryID int    `json:"category_id" form:"category_id"`
	UserID     int    `json:"user_id" form:"user_id"`
	UserName   string `json:"user_name" form:"user_name"`
}

func (u *UpdatePostRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	if err != nil {
		return err
	}

	return nil
}
