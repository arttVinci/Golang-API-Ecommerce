package usecase

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/model/converter"
	"API-Ecommerce-Evermos/internal/repository"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserUsecase interface {
	Register(request model.RegisterUserRequest) (model.UserResponse, error)
	Login(request model.LoginUserRequest) (model.LoginUserResponse, error)
	GetCurrent(userID uint) (model.UserResponse, error)
	Update(userId uint, request model.UpdateUserRequest) (model.UserResponse, error)
}

type UserUsecase struct {
	db        *gorm.DB
	log       *logrus.Logger
	viper     *viper.Viper
	validate  *validator.Validate
	userRepo  repository.IUserRepository
	storeRepo repository.IStoreRepository
}

func NewUserUsecase(
	db *gorm.DB,
	log *logrus.Logger,
	viper *viper.Viper,
	validate *validator.Validate,
	userRepo repository.IUserRepository,
	storeRepo repository.IStoreRepository,
) *UserUsecase {
	{
		return &UserUsecase{
			db:        db,
			log:       log,
			validate:  validate,
			viper:     viper,
			userRepo:  userRepo,
			storeRepo: storeRepo,
		}
	}
}

func (u *UserUsecase) Register(request model.RegisterUserRequest) (model.UserResponse, error) {
	err := u.validate.Struct(request)

	if err != nil {
		u.log.Warnf("Validation failed for register: %v", err)
		return model.UserResponse{}, err
	}

	userCheck, _ := u.userRepo.FindByEmail(request.Email)
	if userCheck.ID != 0 {
		u.log.Warnf("Email already registered: %s", request.Email)
		return model.UserResponse{}, errors.New("email sudah terdaftar")
	}

	userCheck, _ = u.userRepo.FindByNoTelp(request.Notelp)
	if userCheck.ID != 0 {
		u.log.Warnf("Phone number already registered: %s", request.Notelp)
		return model.UserResponse{}, errors.New("nomor telepon sudah terdaftar")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorf("Failed to hash password: %v", err)
		return model.UserResponse{}, err
	}

	newUser := entity.User{
		Nama:     request.Nama,
		Email:    request.Email,
		Notelp:   request.Notelp,
		Password: string(hashedPassword),
		IsAdmin:  false,
	}

	var createdUser entity.User
	tx := u.db.Begin()
	if tx.Error != nil {
		u.log.Errorf("Failed to begin transaction: %v", tx.Error)
		return model.UserResponse{}, tx.Error
	}

	createdUser, err = u.userRepo.Save(newUser)
	if err != nil {
		tx.Rollback()
		u.log.Errorf("Failed to save user in transaction: %v", err)
		return model.UserResponse{}, err
	}

	newStore := entity.Store{
		UserId:   createdUser.ID,
		NamaToko: "Toko " + createdUser.Nama,
		UrlFoto:  "",
	}

	_, err = u.storeRepo.Save(newStore)
	if err != nil {
		tx.Rollback()
		u.log.Errorf("Failed to save store in transaction: %v", err)
		return model.UserResponse{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		u.log.Errorf("Failed to commit transaction: %v", err)
		return model.UserResponse{}, err
	}

	u.log.Infof("User registered successfully: %s (ID: %d)", createdUser.Email, createdUser.ID)

	response := model.UserResponse{
		ID:     createdUser.ID,
		Nama:   createdUser.Nama,
		Email:  createdUser.Email,
		Notelp: createdUser.Notelp,
	}

	return response, nil

}

func (u *UserUsecase) Login(request model.LoginUserRequest) (model.LoginUserResponse, error) {
	err := u.validate.Struct(request)
	if err != nil {
		u.log.Warnf("Validation failed for login: %v", err)
		return model.LoginUserResponse{}, err
	}

	user, err := u.userRepo.FindByEmail(request.Email)
	if err != nil || user.ID == 0 {
		u.log.Warnf("Login attempt failed (email not found): %s", request.Email)
		return model.LoginUserResponse{}, errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		u.log.Warnf("Login attempt failed (password mismatch) for user: %s", request.Email)
		return model.LoginUserResponse{}, errors.New("email atau password salah")
	}

	token, err := u.generateJWT(user)
	if err != nil {
		u.log.Errorf("Failed to generate JWT for user %s: %v", user.Email, err)
		return model.LoginUserResponse{}, err
	}

	userResponse := model.UserResponse{
		ID:      user.ID,
		Nama:    user.Nama,
		Email:   user.Email,
		Notelp:  user.Notelp,
		IsAdmin: user.IsAdmin,
	}

	loginResponse := model.LoginUserResponse{
		User:  userResponse,
		Token: token,
	}

	u.log.Infof("User logged in successfully: %s", user.Email)
	return loginResponse, nil
}

func (u *UserUsecase) GetCurrent(userId uint) (model.UserResponse, error) {
	user, err := u.userRepo.FindById(userId)
	if err != nil {
		u.log.Warnf("User not Found: %d", userId)
		return model.UserResponse{}, errors.New("user tidak ditemukan")
	}

	userResponse := converter.UserToResponse(user)

	return userResponse, nil
}

func (u *UserUsecase) Update(userId uint, request model.UpdateUserRequest) (model.UserResponse, error) {
	user, err := u.userRepo.FindById(userId)

	if err != nil {
		return model.UserResponse{}, errors.New("user tidak ditemukan")
	}

	if request.Nama != nil {
		user.Nama = *request.Nama
	}
	if request.Email != nil {
		user.Email = *request.Email
	}
	if request.Notelp != nil {
		user.Notelp = *request.Notelp
	}
	if request.TanggalLahir != nil {
		user.TanggalLahir = request.TanggalLahir
	}
	if request.Tentang != nil {
		user.Tentang = *request.Tentang
	}
	if request.Pekerjaan != nil {
		user.Pekerjaan = *request.Pekerjaan
	}
	if request.IdProvinsi != nil {
		user.IdProvinsi = *request.IdProvinsi
	}
	if request.IdKota != nil {
		user.IdKota = *request.IdKota
	}

	updatedUser, err := u.userRepo.Update(user)
	if err != nil {
		return model.UserResponse{}, err
	}

	return converter.UserToResponse(updatedUser), nil
}

func (u *UserUsecase) generateJWT(user entity.User) (string, error) {
	jwtSecret := u.viper.GetString("jwt.secret")
	if jwtSecret == "" {
		u.log.Error("JWT_SECRET not found in config")
		return "", errors.New("JWT secret not configured")
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"email":    user.Email,
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token berlaku 3 hari
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
