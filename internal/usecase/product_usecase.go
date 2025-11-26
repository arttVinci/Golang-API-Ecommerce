package usecase

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/model/converter"
	"API-Ecommerce-Evermos/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ProductUsecase struct {
	log          *logrus.Logger
	validate     *validator.Validate
	productRepo  repository.IProductRepository
	storeRepo    repository.IStoreRepository
	categoryRepo repository.ICategoryRepository
}

func NewProductUsecase(
	log *logrus.Logger,
	validate *validator.Validate,
	productRepo repository.IProductRepository,
	storeRepo repository.IStoreRepository,
	categoryRepo repository.ICategoryRepository,
) *ProductUsecase {
	return &ProductUsecase{
		log,
		validate,
		productRepo,
		storeRepo,
		categoryRepo,
	}
}

func (u *ProductUsecase) Create(userId uint, categoryId uint, storeId uint, imagePath string, request model.CreateProductRequest) (model.ProductResponse, error) {
	if err := u.validate.Struct(request); err != nil {
		return model.ProductResponse{}, err
	}

	product := entity.Product{
		TokoID:        storeId,
		CategoryId:    request.CategoryID,
		NamaProduk:    request.NamaProduk,
		Slug:          request.Slug,
		HargaReseller: request.HargaReseller,
		HargaKonsumen: request.HargaKonsumen,
		Stok:          request.Stok,
		Deskripsi:     request.Deskripsi,
	}

	newProduct, err := u.productRepo.Save(product)
	if err != nil {
		return model.ProductResponse{}, err
	}

	if imagePath != "" {
		foto := entity.FotoProduct{
			ProdukID: newProduct.ID,
			Url:      imagePath,
		}
		_, err = u.productRepo.SaveImage(foto)
		if err != nil {
			u.log.Warnf("Gagal menyimpan foto produk ID %d: %v", newProduct.ID, err)
		}
	}

	getStore, err := u.storeRepo.FindById(storeId)

	if err != nil {
		u.log.Warnf("Id Toko tidak di temukan%d: %v", storeId, err)
	}

	getCategory, err := u.categoryRepo.FindById(categoryId)
	if err != nil {
		u.log.Warnf("Id Category tidak di temukan%d: %v", categoryId, err)
	}

	categoryResponse := converter.CategoryToResponse(&getCategory)

	storeResponse := converter.StoreToResponse(getStore)

	productResponse := converter.ProductToResponse(newProduct, storeResponse, categoryResponse)

	return productResponse, nil
}

func (u *ProductUsecase) Search(filter model.ProductFilter) ([]model.ProductResponse, int64, error) {
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	if filter.Page == 0 {
		filter.Page = 1
	}

	products, total, err := u.productRepo.Search(filter)
	if err != nil {
		return nil, 0, err
	}

	var responses []model.ProductResponse
	for _, p := range products {

		responses = append(responses, model.ProductResponse{
			ID:            p.ID,
			NamaProduk:    p.NamaProduk,
			HargaKonsumen: p.HargaKonsumen,
			Stok:          p.Stok,
		})
	}
	return responses, total, nil
}
