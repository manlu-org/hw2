package models

type Reply struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	PostID    uint      `json:"post_id"`
	CreatedAt LocalTime `json:"created_at"`
}
