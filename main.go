package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Password() string {
	err := godotenv.Load("C:/Users/bhuvi/task (1)/task/.env")
	if err != nil {
		fmt.Println(err)
	}
	password := os.Getenv("db")
	return password
}

var (
	username = "root"
	hostname = "127.0.0.1:3306"
)

func dsn(dbName string) string {
	password := Password()
	fmt.Println(password)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)

}

func db() *sql.DB {
	DB, err := sql.Open("mysql", dsn("ecommerce"))
	if err != nil {
		panic("can't able to connect database")
	}
	fmt.Println("ok")
	return DB

}

func main() {
	products()
	stock()
	textfile()
	product_report()
	stock_report()
	sendmail()
	file_reader()

}
