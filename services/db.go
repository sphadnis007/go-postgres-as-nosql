package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-postgres-as-nosql/configs"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	dataSourceName = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		configs.DBUser, configs.DBPassword, configs.DBHost,
		configs.DBPort, configs.DBName, configs.DBSSLMode)
)

func ConnectToDB() (*sql.DB, error) {
	// 1. Get abstract connection structure *sql.DB
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Println("Couldn't Open DB connection: ", err)
		return nil, err
	}

	// 2. Set DB Pool configuration settings
	db.SetMaxOpenConns(configs.DBPoolMaxConns)
	db.SetMaxIdleConns(configs.DBMaxIdleConns)
	db.SetConnMaxLifetime(configs.DBConnLifetime * time.Minute)

	// 3. Ping DB - try to make actual connection with DB
	err = db.Ping()
	if err != nil {
		log.Println("Couldn't Ping DB: ", err)
		return nil, err
	}

	// 4. Drop old tables (deleting old data)
	db.Exec(configs.DropCartTableQuery)

	// 5. Create tables
	_, err = db.Exec(configs.CreateCartTableQuery)
	if err != nil {
		log.Println("Failed to create Cart Table: ", err)
	} else {
		log.Println("Created Cart Table in DB")
	}

	return db, nil
}
