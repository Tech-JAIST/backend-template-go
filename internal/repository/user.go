package repository

import (
	"backend/internal/db/model"
	"context"
)

func (r *Repository) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
