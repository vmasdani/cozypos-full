package main

type ProjectTransactionsView struct {
	Project      Project           `json:"project"`
	Transactions []TransactionView `json:"transactions"`
}

type TransactionView struct {
	Transaction      Transaction           `json:"transaction"`
	ItemTransactions []ItemTransactionView `json:"itemTransactions"`
	TotalPrice       uint                  `json:"totalPrice"`
}

type ItemTransactionView struct {
	Item            Item            `json:"item"`
	ItemTransaction ItemTransaction `json:"itemTransaction"`
}

type ItemStockView struct {
	Item    Item `json:"item"`
	InStock int  `json:"inStock"`
}

type ItemStockInsView struct {
	Item     Item          `json:"item"`
	StockIns []StockInView `json:"stockIns"`
}

type StockInView struct {
	StockIn StockIn `json:"stockIn"`
	Project Project `json:"project"`
}

type ProjectsView struct {
	Projects    []ProjectView `json:"projects"`
	TotalIncome int           `json:"totalIncome"`
}

type ProjectView struct {
	Project                 Project `json:"project"`
	Income                  int     `json:"income"`
	TotalManufacturingPrice uint    `json:"totalManufacturingPrice"`
}
