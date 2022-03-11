package request

type CreateTagRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type CreatePostRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	TagID   int    `json:"tag_id" form:"tag_id"`
}
