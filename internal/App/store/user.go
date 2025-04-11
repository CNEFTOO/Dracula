package store

import (
	"context"

	"github.com/seaung/Dracula/internal/pkg/models"
	"gorm.io/gorm"
)

type UserStore interface {
	GetByID(cxt context.Context, id int64) (*models.Users, error)
	GetByName(cxt context.Context, name string) (*models.Users, error)
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *userStore {
	return &userStore{
		db: db,
	}
}

var _ UserStore = (*userStore)(nil)

func (u *userStore) GetByID(cxt context.Context, id int64) (*models.Users, error) {
	var user models.Users
	if err := u.db.WithContext(cxt).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userStore) GetByName(cxt context.Context, name string) (*models.Users, error) {
	var user models.Users
	if err := u.db.WithContext(cxt).Where("name =?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
