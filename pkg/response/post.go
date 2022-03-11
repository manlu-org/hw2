package response

import "backend-learning/hw2/models"

type PostResp struct {
	List []*models.Post `json:"list"`
	Meta *Meta          `json:"meta"`
}
