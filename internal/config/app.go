package config

import (
	"API-Ecommerce-Evermos/internal/delivery/http"
	"API-Ecommerce-Evermos/internal/delivery/middleware"
	"API-Ecommerce-Evermos/internal/repository"
	"API-Ecommerce-Evermos/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Viper    *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.DB, config.Log)
	storeRepository := repository.NewStoreRepository(config.DB, config.Log)
	addressRepository := repository.NewAddressRepository(config.DB, config.Log)
	categoryRepository := repository.NewCategoryRepository(config.DB, config.Log)
	productRepository := repository.NewProductRepository(config.DB, config.Log)
	transactionRepository := repository.NewTrxRepository(config.Log, config.DB)

	// setup usecases
	userUsecase := usecase.NewUserUsecase(config.DB, config.Log, config.Viper, config.Validate, userRepository, storeRepository)
	storeUsecase := usecase.NewStoreUsecase(config.Log, config.Validate, storeRepository)
	addressUsecase := usecase.NewAddressUsecase(config.Log, config.Validate, addressRepository)
	categoryUsecase := usecase.NewCategoryUsecase(config.Log, config.Validate, categoryRepository)
	productUsecase := usecase.NewProductUsecase(config.Log, config.Validate, productRepository, storeRepository, categoryRepository)
	transactionUsecase := usecase.NewTrxUsecase(config.DB, config.Log, config.Validate, transactionRepository, productRepository)

	// setup controller
	userController := http.NewUserController(userUsecase, config.Log)
	storeController := http.NewStoreController(storeUsecase, config.Log)
	addressController := http.NewAddressController(addressUsecase, config.Log)
	categoryController := http.NewCategoryController(categoryUsecase, config.Log)
	productController := http.NewProductController(productUsecase, storeUsecase, config.Log)
	transactionController := http.NewTrxController(transactionUsecase, config.Log)

	api := config.App.Group("/api/ecommerce")

	api.Post("/register", userController.Register)
	api.Post("/login", userController.Login)
	api.Get("/products", productController.Search)

	protectedApi := api.Group("/")
	protectedApi.Use(middleware.AuthMiddleware(config.Viper))

	protectedApi.Get("/users/current", userController.GetCurrent)
	protectedApi.Put("/users/current", userController.Update)

	protectedApi.Get("/store", storeController.GetMyStore)
	protectedApi.Put("/store", storeController.Update)

	protectedApi.Post("/addresses", addressController.Create)
	protectedApi.Get("/addresses", addressController.List)
	protectedApi.Delete("/addresses/:id", addressController.Delete)

	protectedApi.Post("/products", productController.Create)

	protectedApi.Post("/transactions", transactionController.Create)
	protectedApi.Get("/transactions", transactionController.History)

	adminApi := protectedApi.Group("/admin")
	adminApi.Use(middleware.AdminMiddleware(userRepository))

	adminApi.Post("/categories", categoryController.Create)
	adminApi.Put("/categories/:id", categoryController.Update)
	adminApi.Delete("/categories/:id", categoryController.Delete)

	protectedApi.Get("/categories", categoryController.List)
}
