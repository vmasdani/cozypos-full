package main

import (
  "time"
)

type gormModel struct {
  ID        uint        `gorm:"primary_key" json:"id"`
  CreatedAt time.Time   `json:"createdAt"`
  UpdatedAt time.Time   `json:"updatedAt"`
  DeletedAt *time.Time  `json:"deletedAt"`
}

type ApiKey struct {
  gormModel
  Name                string  `json:"name"`
  Description         string  `json:"description"`
  Price               uint    `json:"price"`
  ManufacturingPrice  uint    `json:"manufacturingPrice"`
}

type Project struct {
  gormModel
  Name      string    `json:"name"`
  StartDate time.Time `json:"startDate"`
}

type Transaction struct {
  gormModel
  Cashier       string  `json:"cashier"`
  PriceIsCustom bool    `json:"priceIsCustom"`
  CustomPrice   uint    `json:"customPrice"`
  ProjectID     uint    `json:"projectId"`
}

type StockIn struct {
  gormModel
  Pic       string  `json:"pic"`
  ItemID    uint    `json:"itemId"`
  ProjectID uint    `json:"projectId"`
  Qty       uint    `json:"qty"`
}

type ItemTransaction struct {
  gormModel
  ItemID        uint  `json:"itemId"`
  TransactionID uint  `json:"transactionId"`
  Qty           uint  `json:"qty"`
}
