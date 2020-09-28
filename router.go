package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Route(r *mux.Router, db *gorm.DB) {
	// Transactions
	r.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		var transactions []Transaction
		All(db, &transactions, w, r)
	}).Methods("GET")

	r.HandleFunc("/transactions/{id}", func(w http.ResponseWriter, r *http.Request) {
		var transaction Transaction
		All(db, &transaction, w, r)
	}).Methods("GET")

	r.HandleFunc("/transactions/{id}", func(w http.ResponseWriter, r *http.Request) {
		var transaction Transaction
		Delete(db, &transaction, w, r)
	}).Methods("DELETE")

	// Items
	r.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		var items []Item
		All(db, &items, w, r)
	}).Methods("GET")

	r.HandleFunc("/items/{id}", func(w http.ResponseWriter, r *http.Request) {
		var item Item
		Get(db, &item, w, r)
	}).Methods("GET")

	r.HandleFunc("/items/{id}", func(w http.ResponseWriter, r *http.Request) {
		var item Item
		Get(db, &item, w, r)
	}).Methods("DELETE")

	r.HandleFunc("/itemstocks", ItemStocks(db))

	// Projects
	r.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		var projects []Project
		All(db, &projects, w, r)
	}).Methods("GET")

	r.HandleFunc("/projects/{id}", func(w http.ResponseWriter, r *http.Request) {
		var project Project
		All(db, &project, w, r)
	}).Methods("GET")

	r.HandleFunc("/projects/{id}", func(w http.ResponseWriter, r *http.Request) {
		var project Project
		Delete(db, &project, w, r)
	}).Methods("GET")

	r.HandleFunc("/projectsview", AllProjectsView(db))
	r.HandleFunc("/projects/{id}/transactions", GetAllProjectTransactions(db))
}
