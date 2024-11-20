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
	host     = "inventory"
	port     = 5434
	user     = "root"
	password = "root"
	dbname   = "inventory"
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

	go handler.OnMessageFromOrderService()
	http.HandleFunc("/inventory/create", handler.Insert)

	log.Println("server started at :8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
