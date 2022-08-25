package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func dsn(dbName string) string {
	var (
		username = "root"
		hostname = "127.0.0.1:3306"
		password = "290400@Ss"
	)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)

}

func db() *sql.DB {
	DB, err := sql.Open("mysql", dsn("ecommerce"))
	if err != nil {
		panic("can't able to connect database")
	}
	return DB

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	products()
	stock()
	textfile()
	product_report()
	stock_report()
	sendmail()
	file_reader()

}
