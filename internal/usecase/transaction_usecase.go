package usecase

import (
	"API-Ecommerce-Evermos/internal/entity"
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/repository"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TrxUsecase struct {
	db          *gorm.DB
	log         *logrus.Logger
	validate    *validator.Validate
	trxRepo     repository.ITrxRepository
	productRepo repository.IProductRepository
}

type ITrxUsecase interface {
	Checkout(userId uint, request model.CreateTransactionRequest) (model.TransactionResponse, error)
	History(userId uint) ([]model.TransactionResponse, error)
}

func NewTrxUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, trxRepo repository.ITrxRepository, productRepo repository.IProductRepository) *TrxUsecase {
	return &TrxUsecase{db, log, validate, trxRepo, productRepo}
}

func (u *TrxUsecase) Checkout(userId uint, request model.CreateTransactionRequest) (model.TransactionResponse, error) {
	if err := u.validate.Struct(request); err != nil {
		return model.TransactionResponse{}, err
	}

	if len(request.Items) == 0 {
		return model.TransactionResponse{}, errors.New("keranjang belanja kosong")
	}

	tx := u.db.Begin()
	if tx.Error != nil {
		return model.TransactionResponse{}, tx.Error
	}

	firstProd, err := u.productRepo.FindById(request.Items[0].ProductID)
	if err != nil {
		u.log.Error("gagal FindById: ", err)
		tx.Rollback()
		return model.TransactionResponse{}, errors.New("produk pertama tidak valid")
	}

	trx := entity.Trx{
		UserId:      userId,
		AlamatId:    uint(request.AlamatPengirimanID),
		TokoId:      firstProd.TokoID,
		MethodBayar: request.MethodBayar,
		KodeInvoice: fmt.Sprintf("INV-%d-%d", userId, time.Now().Unix()),
		HargaTotal:  0,
	}
	newTrx, err := u.trxRepo.CreateTrx(tx, trx)

	if err != nil {
		u.log.Error("Gagal created trx:", err)
		tx.Rollback()
		return model.TransactionResponse{}, err
	}

	var totalHarga int64 = 0
	var detailResponses []model.DetailTransactionResponse

	for _, item := range request.Items {
		product, err := u.productRepo.FindById(item.ProductID)
		if err != nil {
			tx.Rollback()
			return model.TransactionResponse{}, errors.New("produk tidak ditemukan")
		}

		if product.TokoID != newTrx.TokoId {
			tx.Rollback()
			return model.TransactionResponse{}, errors.New("semua barang harus dari toko yang sama")
		}

		if product.Stok < item.Quantity {
			tx.Rollback()
			return model.TransactionResponse{}, fmt.Errorf("stok %s habis/kurang", product.NamaProduk)
		}

		if err := u.productRepo.UpdateStok(tx, product.ID, item.Quantity); err != nil {
			tx.Rollback()
			return model.TransactionResponse{}, err
		}

		logProduct := entity.LogProduct{
			ProdukId:      product.ID,
			TokoID:        product.TokoID,
			CategoryId:    product.CategoryId,
			NamaProduk:    product.NamaProduk,
			Slug:          product.Slug,
			HargaReseller: product.HargaReseller,
			HargaKonsumen: product.HargaKonsumen,
			Deskripsi:     product.Deskripsi,
		}
		newLog, err := u.trxRepo.CreateLogProduct(tx, logProduct)
		if err != nil {
			tx.Rollback()
			return model.TransactionResponse{}, err
		}

		hargaInt, _ := strconv.ParseInt(product.HargaKonsumen, 10, 64)
		subTotal := hargaInt * int64(item.Quantity)
		totalHarga += subTotal

		detail := entity.DetailTrx{
			TrxId:        newTrx.ID,
			TokoId:       product.TokoID,
			LogProductId: newLog.ID,
			Kuantitas:    item.Quantity,
			HargaTotal:   int(subTotal),
		}
		if err := u.trxRepo.CreateDetail(tx, detail); err != nil {
			tx.Rollback()
			return model.TransactionResponse{}, err
		}

		detailResponses = append(detailResponses, model.DetailTransactionResponse{
			ProductName: product.NamaProduk,
			Quantity:    item.Quantity,
			Price:       hargaInt,
			SubTotal:    subTotal,
		})
	}

	newTrx.HargaTotal = totalHarga
	if err := tx.Save(&newTrx).Error; err != nil {
		tx.Rollback()
		return model.TransactionResponse{}, err
	}

	tx.Commit()

	return model.TransactionResponse{
		ID:          newTrx.ID,
		UserID:      newTrx.UserId,
		AlamatID:    uint(newTrx.AlamatId),
		TokoID:      newTrx.TokoId,
		HargaTotal:  newTrx.HargaTotal,
		KodeInvoice: newTrx.KodeInvoice,
		MethodBayar: newTrx.MethodBayar,
		CreatedAt:   newTrx.CreatedAt,
		Details:     detailResponses,
	}, nil
}

func (u *TrxUsecase) History(userId uint) ([]model.TransactionResponse, error) {
	trxes, err := u.trxRepo.FindAllByUserId(userId)
	if err != nil {
		return []model.TransactionResponse{}, err
	}

	var responses []model.TransactionResponse
	for _, trx := range trxes {
		var details []model.DetailTransactionResponse
		for _, d := range trx.DetailTrx {
			details = append(details, model.DetailTransactionResponse{
				ProductName: d.LogProduct.NamaProduk,
				Quantity:    d.Kuantitas,
				Price:       int64(d.HargaTotal) / int64(d.Kuantitas),
				SubTotal:    int64(d.HargaTotal),
			})
		}

		responses = append(responses, model.TransactionResponse{
			ID:          trx.ID,
			UserID:      trx.UserId,
			AlamatID:    uint(trx.AlamatId),
			TokoID:      trx.TokoId,
			HargaTotal:  trx.HargaTotal,
			KodeInvoice: trx.KodeInvoice,
			MethodBayar: trx.MethodBayar,
			CreatedAt:   trx.CreatedAt,
			Details:     details,
		})
	}

	return responses, nil
}
