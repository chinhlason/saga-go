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
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "order"
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
	http.HandleFunc("/order/create", handler.Insert)
	http.HandleFunc("/order/get", handler.Get)
	http.HandleFunc("/order/update", handler.Update)

	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
