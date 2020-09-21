package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func All(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	db.Find(table)
	json.NewEncoder(w).Encode(table)
}

func Get(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("content-type", "application/json")
	db.Where("id = ?", id).Find(table)
	json.NewEncoder(w).Encode(table)
}

func GetAllProjectTransactions(r *mux.Router, db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		id := mux.Vars(r)["id"]
		var project Project
		db.Where("id = ?", id).First(&project)

		var projectTransactionsView ProjectTransactionsView

		transactionViews := []TransactionView{}
		var transactions []Transaction
		db.Order("id desc").Find(&transactions)

		for _, transaction := range transactions {
			var itemTransactions []ItemTransaction
			db.Preload("Item").Where("id = ?", transaction.ID).Find(&itemTransactions)

			totalPrice := uint(0)

			for _, itemTransaction := range itemTransactions {
				totalPrice += itemTransaction.Item.Price * itemTransaction.Qty
			}

			transactionView := TransactionView{
				Transaction:      transaction,
				ItemTransactions: itemTransactions,
				TotalPrice:       totalPrice}

			transactionViews = append(transactionViews, transactionView)
		}

		projectTransactionsView.Project = project
		projectTransactionsView.Transactions = transactionViews

		json.NewEncoder(w).Encode(projectTransactionsView)
	}
}
