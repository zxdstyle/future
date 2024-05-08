package base

import (
	"context"
	"gorm.io/gorm"
)

type Logic[M any] struct {
	db *gorm.DB
}

func NewLogic[M any](db *gorm.DB) *Logic[M] {
	return &Logic[M]{db}
}

func (l *Logic[M]) Show(ctx context.Context, id uint64) (*M, error) {
	var (
		m   M
		err = l.db.WithContext(ctx).Where("id = ?", id).Take(&m).Error
	)
	return &m, err
}
