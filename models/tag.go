package models

type Tag struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt LocalTime `json:"created_at"`
}
