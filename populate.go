package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

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
	// Items
	populateCsv("csv/items.csv", func(items [][]string) {

		for _, item := range items {
			parsedPrice, _ := strconv.ParseUint(item[4], 10, 32)
			parsedManufacturingPrice, _ := strconv.ParseUint(item[5], 10, 32)

			item := Item{
				Uid:                item[1],
				Name:               item[2],
				Description:        item[3],
				Price:              uint(parsedPrice),
				ManufacturingPrice: uint(parsedManufacturingPrice)}
			fmt.Println(item)
			db.Save(&item)
		}
	})

	// Stock In
	populateCsv("csv/itemstockins.csv", func(items [][]string) {
		// for _, item := range items {
		// 	fmt.Println(item)
		// }
	})

	// Item Transaction
	populateCsv("csv/itemtransactions.csv", func(items [][]string) {
		// for _, item := range items {
		// 	fmt.Println(item)
		// }
	})

	// Transaction
	populateCsv("csv/transactions.csv", func(items [][]string) {
		// for _, item := range items {
		// 	fmt.Println(item)
		// }
	})
}
