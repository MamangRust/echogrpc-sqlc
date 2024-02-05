package requests

import "github.com/go-playground/validator/v10"

type CreateCommentRequest struct {
	IdPostComment int    `json:"id_post_comment"`
	Username      string `json:"username"`
	Comment       string `json:"comment"`
}

func (r *CreateCommentRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

type UpdateCommentRequest struct {
	Id            int    `json:"id"`
	IdPostComment int    `json:"id_post_comment"`
	Username      string `json:"username"`
	Comment       string `json:"comment"`
}

func (r *UpdateCommentRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
