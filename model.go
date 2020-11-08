package main

import (
	"time"
)

type GormModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt" ts_type:"string"`
	UpdatedAt time.Time  `json:"updatedAt" ts_type:"string"`
	DeletedAt *time.Time `json:"deletedAt" ts_type:"string"`
}

type ApiKey struct {
	GormModel
	ApiKey string `json:"apiKey"`
}

type Item struct {
	GormModel
	Uid                string            `json:"uid"`
	Name               string            `json:"name"`
	Description        string            `json:"description"`
	Price              uint              `json:"price"`
	ManufacturingPrice uint              `json:"manufacturingPrice"`
	ItemTransactions   []ItemTransaction `json:"itemTransactions" ts_type:"ItemTransaction[] | null"`
	StockIns           []StockIn         `json:"stockIns" ts_type:"StockIn[] | null"`
}

type Project struct {
	GormModel
	Uid          string        `json:"uid"`
	Name         string        `json:"name"`
	StartDate    time.Time     `json:"startDate" ts_type:"string"`
	Transactions []Transaction `json:"transactions" ts_type:"Transaction[] | null"`
}

type Transaction struct {
	GormModel
	Uid              string            `json:"uid"`
	Cashier          string            `json:"cashier"`
	PriceIsCustom    bool              `json:"priceIsCustom"`
	CustomPrice      uint              `json:"customPrice"`
	ProjectID        uint              `json:"projectId"`
	ItemTransactions []ItemTransaction `json:"itemTransactions" ts_type:"ItemTransaction[] | null"`
}

type StockIn struct {
	GormModel
	Uid       string `json:"uid"`
	Pic       string `json:"pic"`
	ItemID    uint   `json:"itemId"`
	ProjectID uint   `json:"projectId"`
	Qty       uint   `json:"qty"`
}

type ItemTransaction struct {
	GormModel
	Uid           string `json:"uid"`
	ItemID        uint   `json:"itemId"`
	TransactionID uint   `json:"transactionId"`
	Qty           uint   `json:"qty"`
	Item          Item   `json:"item"`
}
