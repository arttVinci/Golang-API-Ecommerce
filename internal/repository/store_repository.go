package repository

import (
	"API-Ecommerce-Evermos/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IStoreRepository interface {
	Save(store entity.Store) (entity.Store, error)
	FindByUserId(userId uint) (entity.Store, error)
	FindById(storeId uint) (entity.Store, error)
	Update(store entity.Store) (entity.Store, error)
}

type StoreRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewStoreRepository(db *gorm.DB, log *logrus.Logger) *StoreRepository {
	return &StoreRepository{db, log}
}

func (r StoreRepository) Save(store entity.Store) (entity.Store, error) {
	err := r.db.Create(&store).Error

	if err != nil {
		r.log.Errorf("Error saving store: %v", err)
		return store, err
	}

	r.log.Tracef("Store saved with id: %d", store.ID)
	return store, nil
}

func (r *StoreRepository) FindByUserId(userId uint) (entity.Store, error) {
	var store entity.Store

	err := r.db.Where("user_id = ?", userId).First(&store).Error
	if err != nil {
		return store, err
	}
	return store, nil
}

func (r *StoreRepository) FindById(storeId uint) (entity.Store, error) {
	var store entity.Store

	err := r.db.Where("id = ?", storeId).First(&store).Error

	if err != nil {
		return store, err
	}
	return store, nil

}

func (r *StoreRepository) Update(store entity.Store) (entity.Store, error) {
	err := r.db.Save(&store).Error
	if err != nil {
		r.log.Errorf("Error updating store: %v", err)
		return store, err
	}
	return store, nil
}
