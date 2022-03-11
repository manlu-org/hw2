package request

type CreateReplyRequest struct {
	Content string `json:"content" form:"content" validate:"required"`
	PostID  int    `json:"post_id" form:"post_id" validate:"required"`
}
