package service

import (
	"backend-learning/hw2/global"
	"backend-learning/hw2/models"
)

var Post = new(postService)

type postService struct {
}

func (pp postService) CreatePost(p *models.Post) (*models.Post, error) {
	if err := global.DB.Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (pp postService) GetPost(id int) (*models.Post, error) {
	post := new(models.Post)
	if err := global.DB.Model(&models.Post{}).
		Preload("Tags").
		Preload("Replies").
		First(post, id).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (pp postService) GetAllPosts(page, pageSize int) ([]*models.Post, int64, error) {
	posts := make([]*models.Post, 0)
	var total int64
	query := global.DB.Model(&models.Post{})
	query.Count(&total)

	if page > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	} else {
		query = query.Limit(pageSize)
	}

	if err := query.Order("id desc").Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (pp postService) ReplyPost(reply *models.Reply) error {
	if err := global.DB.Create(reply).Error; err != nil {
		return err
	}
	return nil
}

func (pp postService) DelPost(id int) error {
	if err := global.DB.Delete(&models.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}
