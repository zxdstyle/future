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

func (l *Logic[M]) Create(ctx context.Context, m *M) error {
	return l.db.WithContext(ctx).Create(m).Error
}

func (l *Logic[M]) Update(ctx context.Context, id uint64, m *M) error {
	return l.db.WithContext(ctx).Where("id = ?", id).Updates(&m).Error
}

func (l *Logic[M]) Destroy(ctx context.Context, id uint64) error {
	var m M
	return l.db.WithContext(ctx).Where("id = ?", id).Delete(&m).Error
}
