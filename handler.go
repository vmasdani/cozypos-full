package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"gorm.io/gorm"
)

func Ts() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		converter := typescriptify.New()

		structs := []interface{}{
			ApiKey{},
			Item{},
			Project{},
			Transaction{},
			StockIn{},
			ItemTransaction{}}

		for _, myStruct := range structs {
			converter.Add(myStruct)
		}

		converter.ConvertToFile("./newmodel.d.ts")
	}
}

func Login(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func All(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	db.Find(table)
	json.NewEncoder(w).Encode(table)
}

func Get(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// fmt.Println("ID:", id)
	w.Header().Set("content-type", "application/json")
	db.Where("id = ?", id).First(table)
	json.NewEncoder(w).Encode(table)
}

func Post(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewDecoder(r.Body).Decode(table)
	db.Save(table)
	json.NewEncoder(w).Encode(table)
	w.WriteHeader(http.StatusCreated)
}

func Delete(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("content-type", "application/json")
	db.Where("id = ?", id).Delete(table)
}

func GetAllProjectTransactions(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		id := mux.Vars(r)["id"]
		var project Project
		db.Where("id = ?", id).First(&project)

		var projectTransactionsView ProjectTransactionsView

		transactionViews := []TransactionView{}
		var transactions []Transaction
		db.Where("project_id = ?", project.ID).Order("id desc").Find(&transactions)

		for _, transaction := range transactions {
			var itemTransactions []ItemTransaction
			db.Preload("Item").Where("id = ?", transaction.ID).Find(&itemTransactions)

			totalPrice := uint(0)

			for _, itemTransaction := range itemTransactions {
				totalPrice += itemTransaction.Item.Price * itemTransaction.Qty
			}

			transactionView := TransactionView{
				Transaction: transaction,
				// ItemTransactions: itemTransactions,
				TotalPrice: totalPrice}

			for _, itemTransaction := range itemTransactions {
				itemTransactionView := ItemTransactionView{
					Item:            itemTransaction.Item,
					ItemTransaction: itemTransaction}

				transactionView.ItemTransactions = append(transactionView.ItemTransactions, itemTransactionView)
			}

			transactionViews = append(transactionViews, transactionView)
		}

		projectTransactionsView.Project = project
		projectTransactionsView.Transactions = transactionViews

		json.NewEncoder(w).Encode(projectTransactionsView)
	}
}

func ItemStocks(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var items []Item
		db.Preload("ItemTransactions").Preload("StockIns").Find(&items)

		itemStockViews := []ItemStockView{}

		// fmt.Println("Fetched items: ", len(items))

		for _, item := range items {
			totalQty := 0

			for _, stockIn := range item.StockIns {
				totalQty += int(stockIn.Qty)
			}

			for _, itemTransaction := range item.ItemTransactions {
				totalQty -= int(itemTransaction.Qty)
			}

			itemStockViews = append(itemStockViews, ItemStockView{
				Item:    item,
				InStock: totalQty})
		}

		json.NewEncoder(w).Encode(&itemStockViews)
	}
}

func AllProjectsView(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		projectsView := ProjectsView{}
		var projects []Project
		db.Preload("Transactions.ItemTransactions.Item").Find(&projects)

		for _, project := range projects {
			projectIncome := 0

			for _, transaction := range project.Transactions {
				transactionTotal := uint(0)

				// fmt.Println("Transaction ID", transaction.ID)

				if transaction.PriceIsCustom {
					// fmt.Println("Is custom", transaction.CustomPrice)
					transactionTotal += transaction.CustomPrice
				} else {
					totalTransactionPrice := uint(0)

					for _, itemTransaction := range transaction.ItemTransactions {
						totalTransactionPrice += itemTransaction.Qty * itemTransaction.Item.Price
					}

					transactionTotal += totalTransactionPrice
				}

				projectIncome += int(transactionTotal)
			}

			projectsView.Projects = append(projectsView.Projects, ProjectView{
				Project: project,
				Income:  projectIncome})
		}

		json.NewEncoder(w).Encode(projectsView)
	}
}

func ViewTransaction(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var transaction Transaction
		db.Preload("ItemTransactions.Item").Where("id = ?", id).First(&transaction)

		var transactionView TransactionView

		transactionView.Transaction = transaction

		for _, itemTransaction := range transaction.ItemTransactions {
			transactionView.ItemTransactions = append(transactionView.ItemTransactions, ItemTransactionView{
				Item:            itemTransaction.Item,
				ItemTransaction: itemTransaction})
		}

		json.NewEncoder(w).Encode(transactionView)
	}
}
