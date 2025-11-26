package model

import "time"

type TransactionItemRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,min=1"`
}

type CreateTransactionRequest struct {
	AlamatPengirimanID uint                     `json:"alamat_id" validate:"required"`
	MethodBayar        string                   `json:"method_bayar" validate:"required"`
	Items              []TransactionItemRequest `json:"items" validate:"required,min=1,dive"`
}

type TransactionResponse struct {
	ID          uint                        `json:"id"`
	UserID      uint                        `json:"user_id"`
	AlamatID    uint                        `json:"alamat_id"`
	TokoID      uint                        `json:"toko_id"`
	HargaTotal  int64                       `json:"harga_total"`
	KodeInvoice string                      `json:"kode_invoice"`
	MethodBayar string                      `json:"method_bayar"`
	CreatedAt   time.Time                   `json:"created_at"`
	Details     []DetailTransactionResponse `json:"details"`
}

type DetailTransactionResponse struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int64  `json:"price"`
	SubTotal    int64  `json:"sub_total"`
}
