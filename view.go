package main

type ProjectTransactionsView struct {
	Project      Project           `json:"project"`
	Transactions []TransactionView `json:"transactions"`
}

type TransactionView struct {
	Transaction      Transaction       `json:"transaction"`
	ItemTransactions []ItemTransaction `json:"itemTransactions"`
	TotalPrice       uint              `json:"totalPrice"`
}
