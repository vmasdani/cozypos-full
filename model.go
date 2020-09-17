package main

import (
	"time"
)

type GormModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type ApiKey struct {
	GormModel
	Name               string `json:"name"`
	Description        string `json:"description"`
	Price              uint   `json:"price"`
	ManufacturingPrice uint   `json:"manufacturingPrice"`
}

type Project struct {
	GormModel
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
}

type Transaction struct {
	GormModel
	Cashier       string `json:"cashier"`
	PriceIsCustom bool   `json:"priceIsCustom"`
	CustomPrice   uint   `json:"customPrice"`
	ProjectID     uint   `json:"projectId"`
}

type StockIn struct {
	GormModel
	Pic       string `json:"pic"`
	ItemID    uint   `json:"itemId"`
	ProjectID uint   `json:"projectId"`
	Qty       uint   `json:"qty"`
}

type ItemTransaction struct {
	GormModel
	ItemID        uint `json:"itemId"`
	TransactionID uint `json:"transactionId"`
	Qty           uint `json:"qty"`
}
