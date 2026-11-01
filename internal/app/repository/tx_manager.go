package repository

import (
	"context"

	"gorm.io/gorm"
)

type TxManager interface {
	Run(ctx context.Context, f func(*gorm.DB) error) error
}

type txManager struct {
	db *gorm.DB
}

func NewTxManager(db *gorm.DB) TxManager {
	return &txManager{db}
}

func (t *txManager) Run(ctx context.Context, f func(*gorm.DB) error) error {
	tx := t.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := f(tx)
	if err != nil {
		return err
	} else {
		return tx.Commit().Error
	}
}
