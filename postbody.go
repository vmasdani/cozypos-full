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
