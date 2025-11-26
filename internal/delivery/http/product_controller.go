package http

import (
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/usecase"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	usecase      *usecase.ProductUsecase
	storeUsecase *usecase.StoreUsecase
	log          *logrus.Logger
}

func NewProductController(usecase *usecase.ProductUsecase, storeUsecase *usecase.StoreUsecase, log *logrus.Logger) *ProductController {
	return &ProductController{usecase, storeUsecase, log}
}

func (c *ProductController) Create(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)

	categoryID, _ := strconv.Atoi(ctx.FormValue("category_id"))
	stok, _ := strconv.Atoi(ctx.FormValue("stok"))

	request := model.CreateProductRequest{
		CategoryID:    uint(categoryID),
		NamaProduk:    ctx.FormValue("nama_produk"),
		Slug:          ctx.FormValue("slug"),
		HargaReseller: ctx.FormValue("harga_reseller"),
		HargaKonsumen: ctx.FormValue("harga_konsumen"),
		Stok:          stok,
		Deskripsi:     ctx.FormValue("deskripsi"),
	}

	imagePath := ""
	file, err := ctx.FormFile("image")
	if err == nil {
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		imagePath = fmt.Sprintf("public/uploads/%s", filename)

		errSave := ctx.SaveFile(file, "./"+imagePath)
		if errSave != nil {
			c.log.Errorf("Gagal save file: %v", errSave)
		}
	}

	myStore, err := c.storeUsecase.GetMyStore(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Anda belum memiliki toko")
	}

	response, err := c.usecase.Create(userId, uint(categoryID), myStore.ID, imagePath, request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": response})
}

func (c *ProductController) Search(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	categoryId := ctx.QueryInt("category_id", 0)
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	filter := model.ProductFilter{
		Search:     search,
		CategoryID: categoryId,
		Page:       page,
		Limit:      limit,
	}

	responses, total, err := c.usecase.Search(filter)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   responses,
		"meta": fiber.Map{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}
