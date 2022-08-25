package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func whileerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var id = 1

func products() {
	data := db()
	defer data.Close()
	excel1, err := excelize.OpenFile("product_details.xlsx")
	whileerr(err)
	sh1, err := excel1.GetRows("Sheet1")
	whileerr(err)
	for _, rows := range sh1 {
		result, err := data.Query("SELECT product_id from product_stock")
		whileerr(err)
		var product_id int
		var list_of_id []int
		for result.Next() {
			result.Scan(&product_id)
			list_of_id = append(list_of_id, product_id)
		}
		val, _ := strconv.Atoi(rows[0])
		var ispresent bool = true
		for _, v := range list_of_id {
			if val == v {
				ispresent = false
				break
			} else {
				ispresent = true
			}
		}
		if ispresent {
			res, err := data.Query("insert into product(product_id, product_name, product_price) values (?,?,?)", rows[0], rows[1], rows[2])
			whileerr(err)
			defer res.Close()
		}
	}
}

func stock() {
	data := db()
	defer data.Close()
	excel2, err := excelize.OpenFile("stock_details.xlsx")
	whileerr(err)
	sh2, err := excel2.GetRows("Sheet1")
	whileerr(err)
	for _, row := range sh2 {
		result, _ := data.Query("SELECT product_id from product_stock")
		var product_id int
		var list_of_id []int
		for result.Next() {
			result.Scan(&product_id)
			list_of_id = append(list_of_id, product_id)
		}
		val, _ := strconv.Atoi(row[0])
		var ispresent bool = true
		for _, v := range list_of_id {
			if val == v {
				ispresent = false
				break
			} else {
				ispresent = true
			}
		}
		if ispresent {
			res, err := data.Query("insert into product_stock(product_id, batch_id, stock_qty) values (?,?,?)", row[0], id, row[1])
			whileerr(err)
			defer res.Close()

		}
	}
	fmt.Println("completed insert")

}
