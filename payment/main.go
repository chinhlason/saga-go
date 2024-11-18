package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"order/pkg"
)

var (
	host     = "localhost"
	port     = 5433
	user     = "root"
	password = "root"
	dbname   = "payment"
	driver   = "postgres"
	sslMode  = "disable"
)

func main() {
	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslMode)
	db, err := sql.Open(driver, dbConnStr)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	handler := pkg.NewHandler(pkg.NewRepository(db))
	http.HandleFunc("/payment/create", handler.Insert)
	http.HandleFunc("/payment/get", handler.Get)
	http.HandleFunc("/payment/update", handler.Update)

	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
