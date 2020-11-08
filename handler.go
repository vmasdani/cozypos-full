package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Ts() {
	converter := typescriptify.New()
	converter.CreateInterface = true

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

func Login(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginPostBody LoginPostBody
		json.NewDecoder(r.Body).Decode(&loginPostBody)

		err := godotenv.Load()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error loading .env!")
			log.Fatal("Error loading .env file.")
			return
		}

		password := os.Getenv("PASSWORD")
		// fmt.Println(loginPostBody.Password, password)
		passwordNotMatchErr := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginPostBody.Password))

		if passwordNotMatchErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Password do not match!")
			return
		}

		encodedUsername := base64.StdEncoding.EncodeToString([]byte(loginPostBody.Username))

		var apiKey ApiKey
		if err := db.Where("api_key like ?", fmt.Sprintf("%s%%", encodedUsername)).First(&apiKey).Error; err != nil {
			fmt.Println("API key record not found! Creating..")
		} else {
			db.Delete(&apiKey)
		}

		timestamp, _ := bcrypt.GenerateFromPassword([]byte(time.Now().String()), bcrypt.DefaultCost)
		// fmt.Println("timestamp", string(timestamp))

		apiKeyString := fmt.Sprintf("%s:%s", encodedUsername, timestamp)

		db.Save(&ApiKey{
			ApiKey: apiKeyString})

		fmt.Fprintf(w, "%s", apiKeyString)
	}
}

func GeneratePassword() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// var passwordPostBody PasswordPostBody
		// json.NewDecoder(r.Body).Decode(&passwordPostBody)
		password := r.URL.Query()["password"]

		if len(password) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "No password query param provided.")
			return
		}

		generatedPassword, _ := bcrypt.GenerateFromPassword([]byte(password[0]), bcrypt.DefaultCost)
		fmt.Fprintf(w, "'%s'", string(generatedPassword))
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
			db.Preload("Item").Where("transaction_id = ?", transaction.ID).Find(&itemTransactions)

			totalPrice := uint(0)

			for _, itemTransaction := range itemTransactions {
				totalPrice += itemTransaction.Item.Price * itemTransaction.Qty
			}

			transactionView := TransactionView{
				Transaction:      transaction,
				ItemTransactions: []ItemTransactionView{},
				TotalPrice:       totalPrice}

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
		db.
			Preload("ItemTransactions").
			Preload("StockIns").
			Order("id desc").
			Find(&items)

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

func ItemSearch(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var items []Item

		itemName := r.URL.Query()["name"]

		if len(itemName) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "No name parameter provided.")
			return
		} else {
			db.Limit(5).Where("name like ?", fmt.Sprintf("%%%s%%", itemName[0])).Find(&items)

			itemStockViews := []ItemStockView{}

			for _, item := range items {
				var stockIns []StockIn
				var itemTransactions []ItemTransaction

				totalStock := 0

				db.Where("item_id = ?", item.ID).Find(&stockIns)
				db.Where("item_id = ?", item.ID).Find(&itemTransactions)

				for _, stockIn := range stockIns {
					totalStock += int(stockIn.Qty)
				}

				for _, itemTransaction := range itemTransactions {
					totalStock -= int(itemTransaction.Qty)
				}

				itemStockViews = append(itemStockViews, ItemStockView{
					Item:    item,
					InStock: totalStock})
			}

			json.NewEncoder(w).Encode(&itemStockViews)
		}
	}
}

func SaveTransaction(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var transactionPostBody TransactionPostBody
		json.NewDecoder(r.Body).Decode(&transactionPostBody)

		db.Save(&transactionPostBody.Transaction)

		for _, itemTransactionView := range transactionPostBody.ItemTransactions {
			itemTransactionView.ItemTransaction.TransactionID = transactionPostBody.Transaction.ID
			db.Save(&itemTransactionView.ItemTransaction)
		}

		for _, itemTransactionID := range transactionPostBody.ItemTransactionDeleteIds {
			// var itemTransaction ItemTransaction
			// fmt.Println("Item Transaction ID:", itemTransactionID)
			// db.Where("id = ?", itemTransactionID).Delete(&itemTransaction)
			db.Delete(&ItemTransaction{}, itemTransactionID)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(transactionPostBody.Transaction)
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
		transactionView.ItemTransactions = []ItemTransactionView{}

		for _, itemTransaction := range transaction.ItemTransactions {
			transactionView.ItemTransactions = append(transactionView.ItemTransactions, ItemTransactionView{
				Item:            itemTransaction.Item,
				ItemTransaction: itemTransaction})
		}

		json.NewEncoder(w).Encode(transactionView)
	}
}

func ItemStockIns(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		w.Header().Set("content-type", "application/json")
		var item Item
		db.Where("id = ?", id).First(&item)

		stockInViews := []StockInView{}

		var stockIns []StockIn
		db.Where("item_id = ?", item.ID).Find(&stockIns)

		for _, stockIn := range stockIns {
			var project Project
			db.Where("id = ?", stockIn.ProjectID).Find(&project)

			stockInViews = append(stockInViews, StockInView{
				StockIn: stockIn,
				Project: project})
		}

		json.NewEncoder(w).Encode(&ItemStockInsView{
			Item:     item,
			StockIns: stockInViews})
	}
}

func ItemStockInsAdd(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// id := mux.Vars(r)["id"]
		w.Header().Set("content-type", "application/json")
		var stockIn StockIn
		json.NewDecoder(r.Body).Decode(&stockIn)
		db.Save(&stockIn)
		w.WriteHeader(http.StatusCreated)
	}
}

func ItemSave(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var itemPostBody ItemPostBody
		json.NewDecoder(r.Body).Decode(&itemPostBody)

		db.Save(&itemPostBody.Item)

		newStockIn := StockIn{
			ItemID: itemPostBody.Item.ID,
			Qty:    itemPostBody.InitialStockQty,
		}

		db.Save(&newStockIn)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&itemPostBody.Item)
	}
}
