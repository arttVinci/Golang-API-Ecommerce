package model

type ProductFilter struct {
	Page       int
	Limit      int
	Search     string
	CategoryID int
}
