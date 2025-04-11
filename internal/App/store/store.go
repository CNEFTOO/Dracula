package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once *sync.Once
	S    database
)

type database struct {
	db *gorm.DB
}

type Store interface {
	DB() *gorm.DB
	Users() UserStore
}

var _ Store = (*database)(nil)

func NewStore(db *gorm.DB) *database {
	once = &sync.Once{}
	once.Do(func() {
		S = database{
			db: db,
		}
	})
	return &S
}

func (s *database) DB() *gorm.DB {
	return s.db
}

func (s *database) Users() UserStore {
	return NewUserStore(s.db)
}
