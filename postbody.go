package main

type LoginPostBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordPostBody struct {
	Password string `json:"password"`
}

type TransactionPostBody struct {
	Transaction              Transaction           `json:"transaction"`
	ItemTransactions         []ItemTransactionView `json:"itemTransactions"`
	ItemTransactionDeleteIds []uint                `json:"itemTransactionDeleteIds"`
}

type ItemPostBody struct {
	InitialStockQty  uint    `json:"initialStockQty"`
	Item             Item    `json:"item"`
	Project          Project `json:"project"`
	WithInitialStock bool    `json:"withInitialStock"`
}
