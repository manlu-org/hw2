package request

type CreateTagRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}
