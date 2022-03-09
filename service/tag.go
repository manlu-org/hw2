package service

import (
	"backend-learning/hw2/global"
	"backend-learning/hw2/models"
)

var Tag = new(tagService)

type tagService struct {
}

func (t *tagService) CreateTag(tag *models.Tag) (*models.Tag, error) {
	if err := global.DB.Create(tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}
