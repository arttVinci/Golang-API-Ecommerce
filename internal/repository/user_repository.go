package repository

import (
	"API-Ecommerce-Evermos/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Save(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByNoTelp(notelp string) (entity.User, error)
	FindById(userId uint) (entity.User, error)
}

type UserRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewUserRepository(db *gorm.DB, log *logrus.Logger) *UserRepository {
	return &UserRepository{db, log}
}

func (r UserRepository) Save(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		r.log.Errorf("Error saving user: %v", err)
		return user, err
	}

	r.log.Tracef("User saved with id: $d", user.ID)
	return user, nil
}

func (r *UserRepository) Update(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		r.log.Errorf("Error updating user: %v", err)
		return user, err
	}
	return user, nil
}

func (r UserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		r.log.Errorf("Error finding email: %v", err)
		return user, err
	}

	return user, nil
}

func (r UserRepository) FindByNoTelp(notelp string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("no_telp = ?", notelp).Find(&user).Error

	if err != nil {
		r.log.Errorf("Error finding No Telp: %v", err)
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindById(id uint) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
