package request

type PageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type SignInRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
