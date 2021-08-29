package main

import (
	"log"

	"github.com/go-postgres-as-nosql/services"
)

func main() {

	// 1. Connect to DB
	db, err := services.ConnectToDB()
	if err != nil {
		return
	} else {
		log.Println("Connected to DB")
	}
	defer db.Close()

	// 2. Start service to handle products
	pOps := services.NewProductOps(db)

	// 3. Start REST server and expose URLs
	services.StartRESTServer(pOps)

	select {}
}
