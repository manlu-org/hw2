package models

type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt LocalTime `json:"created_at"`

	Tags    []*Tag   `gorm:"many2many:tag_posts" json:"tags,omitempty"`
	Replies []*Reply `json:"replies,omitempty"`
}

type PostTag struct {
	PostID uint
	TagID  uint
}
