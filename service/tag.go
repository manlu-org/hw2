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

func (t *tagService) GetAllTags() ([]*models.Tag, error) {
	tags := make([]*models.Tag, 0)

	if err := global.DB.Model(&models.Tag{}).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t tagService) GetTag(id int) (*models.Tag, error) {
	tag := new(models.Tag)
	if err := global.DB.Model(&models.Tag{}).First(tag, id).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

func (t tagService) GetPostByTagID(id int) ([]*models.Post, error) {
	tag := new(models.Tag)

	if err := global.DB.Model(&models.Tag{}).Preload("Posts").First(tag, id).Error; err != nil {
		return nil, err
	}

	return tag.Posts, nil
}
