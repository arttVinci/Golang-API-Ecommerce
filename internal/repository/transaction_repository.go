package repository

import (
	"API-Ecommerce-Evermos/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ITrxRepository interface {
	CreateTrx(tx *gorm.DB, trx entity.Trx) (entity.Trx, error)
	CreateDetail(tx *gorm.DB, detail entity.DetailTrx) error
	CreateLogProduct(tx *gorm.DB, log entity.LogProduct) (entity.LogProduct, error)
	FindAllByUserId(userId uint) ([]entity.Trx, error)
}

type TrxRepository struct {
	log *logrus.Logger
	db  *gorm.DB
}

func NewTrxRepository(log *logrus.Logger, db *gorm.DB) *TrxRepository {
	return &TrxRepository{log, db}
}

func (r *TrxRepository) CreateTrx(tx *gorm.DB, trx entity.Trx) (entity.Trx, error) {
	err := tx.Create(&trx).Error
	return trx, err
}

func (r *TrxRepository) CreateDetail(tx *gorm.DB, detail entity.DetailTrx) error {
	return tx.Create(&detail).Error
}

func (r *TrxRepository) CreateLogProduct(tx *gorm.DB, log entity.LogProduct) (entity.LogProduct, error) {
	err := tx.Create(&log).Error
	return log, err
}

func (r *TrxRepository) FindAllByUserId(userId uint) ([]entity.Trx, error) {
	var trxes []entity.Trx
	err := r.db.Preload("DetailTrx.LogProduct").
		Where("user_id = ?", userId).
		Order("created_at desc").
		Find(&trxes).Error
	return trxes, err
}
