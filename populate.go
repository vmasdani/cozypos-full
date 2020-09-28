package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

func populateCsv(path string, parse func(arr [][]string)) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	csvString := string(content)
	// fmt.Println(itemsCsv)

	r := csv.NewReader(strings.NewReader(csvString))

	var records [][]string

	r.Read()
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, record)
	}

	parse(records)
}

func populate(db *gorm.DB) {
	// Project
	timeParsed, _ := time.Parse(time.RFC3339, "2020-02-23T22:08:41+00:00")

	project := Project{
		Name:      "CF14",
		StartDate: timeParsed}
	db.Save(&project)

	fmt.Println(project)

	// Items
	populateCsv("csv/items.csv", func(items [][]string) {

		for _, item := range items {
			parsedPrice, _ := strconv.ParseUint(item[4], 10, 32)
			parsedManufacturingPrice, _ := strconv.ParseUint(item[5], 10, 32)

			item := Item{
				Uid:                fmt.Sprintf("item-%s", item[0]),
				Name:               item[2],
				Description:        item[3],
				Price:              uint(parsedPrice),
				ManufacturingPrice: uint(parsedManufacturingPrice)}
			// fmt.Println(item)
			db.Save(&item)
		}
	})

	// Transaction
	populateCsv("csv/transactions.csv", func(transactions [][]string) {
		for _, transaction := range transactions {
			customPrice, _ := strconv.ParseUint(transaction[2], 10, 32)

			priceIsCustom := customPrice > 0

			transaction := Transaction{
				Uid:           fmt.Sprintf("transaction-%s", transaction[0]),
				CustomPrice:   uint(customPrice),
				Cashier:       transaction[3],
				ProjectID:     project.ID,
				PriceIsCustom: priceIsCustom}

			db.Save(&transaction)
		}
	})

	// Stock In
	populateCsv("csv/itemstockins.csv", func(itemStockIns [][]string) {
		for _, itemStockIn := range itemStockIns {
			itemId, _ := strconv.ParseUint(itemStockIn[3], 10, 32)
			qty, _ := strconv.ParseUint(itemStockIn[4], 10, 32)

			var foundItem Item

			if err := db.Where("uid = ?", fmt.Sprintf("item-%d", itemId)).First(&foundItem).Error; err != nil {
				// fmt.Println(foundItem)
				fmt.Printf("Item not found!, %d", itemId)
			} else {
				// fmt.Println("Item found!")
				// fmt.Println(foundItem)
				// fmt.Printf("Qty: %d", qty)

				itemStockIn := StockIn{
					Uid:       itemStockIn[1],
					Pic:       itemStockIn[2],
					ItemID:    uint(foundItem.ID),
					ProjectID: project.ID,
					Qty:       uint(qty)}

				db.Save(&itemStockIn)
			}
		}
	})

	// Item Transaction
	populateCsv("csv/itemtransactions.csv", func(itemTransactions [][]string) {
		for _, itemTransaction := range itemTransactions {
			itemId := itemTransaction[2]
			transactionId := itemTransaction[3]
			qty, _ := strconv.ParseUint(itemTransaction[4], 10, 32)

			var foundItem Item
			var foundTransaction Transaction

			itemUid := fmt.Sprintf("item-%s", itemId)
			transactionUid := fmt.Sprintf("transaction-%s", transactionId)

			itemErr := db.Where("uid = ?", itemUid).First(&foundItem).Error
			transactionErr := db.Where("uid = ?", transactionUid).First(&foundTransaction).Error

			var foundItemID uint
			var foundTransactionID uint

			if itemErr != nil {
				fmt.Println("Item not found.")
			} else {
				foundItemID = foundItem.ID
			}

			if transactionErr != nil {
				fmt.Println("Transaction not found.")
			} else {
				foundTransactionID = foundTransaction.ID
			}

			fmt.Printf("Trans id %s item id %s\n", transactionUid, itemUid)

			itemTransaction := ItemTransaction{
				Uid:           fmt.Sprintf("itemtransaction-%s", itemTransaction[0]),
				ItemID:        foundItemID,
				TransactionID: foundTransactionID,
				Qty:           uint(qty)}

			db.Save(&itemTransaction)
		}
	})
}
