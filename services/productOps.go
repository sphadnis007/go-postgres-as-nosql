package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-postgres-as-nosql/schema"
	"github.com/google/uuid"
)

type ProductOpsI interface {
	Get(id string) (*[]schema.Product, bool)
	GetAllCarts() ([]string, error)
	AddFakeData(numOfObj int) (string, error)
	// Update()
	Delete(cartID, productID string) error
}

type ProductOps struct {
	DB *sql.DB
}

func NewProductOps(db *sql.DB) ProductOpsI {
	return &ProductOps{DB: db}
}

func (p *ProductOps) Get(id string) (*[]schema.Product, bool) {
	var product []schema.Product
	var dbScan []byte

	if err := p.DB.QueryRow("SELECT product_details FROM cart WHERE cart_id = $1",
		id).Scan(&dbScan); err != nil {
		if err == sql.ErrNoRows {
			log.Println("No such cart is available")
			return nil, true
		}
		log.Println("Failed to get products: ", err)
		return nil, false
	}

	if err := json.Unmarshal(dbScan, &product); err != nil {
		log.Println("Failed: ", err)
		return nil, false
	}

	return &product, false
}

func (p *ProductOps) GetAllCarts() ([]string, error) {
	var cartIDs []string
	var cartID string

	rows, err := p.DB.Query("SELECT cart_id FROM cart")
	if err != nil {
		log.Println("Failed to get carts: ", err)
		return cartIDs, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&cartID); err != nil {
			log.Println("Failed: ", err)
			return cartIDs, err
		}

		cartIDs = append(cartIDs, cartID)
	}

	if rows.Err() != nil {
		// we may have partial result set
		// so, not thowing error to user - just logging it
		log.Println("Error occured while scaning rows: ", err)
	}

	return cartIDs, nil
}

func (p *ProductOps) AddFakeData(numOfObj int) (string, error) {
	var product schema.Product
	products := make([]schema.Product, numOfObj)
	for i := 0; i < numOfObj; i++ {
		gofakeit.Struct(&product)
		products[i] = product
	}

	productsInBytes, err := json.Marshal(products)
	if err != nil {
		log.Println("Failed to generate fake data: ", err)
		return "", err
	}

	cartID := uuid.NewString()

	_, err = p.DB.Exec("INSERT INTO cart VALUES ($1, $2)", cartID, productsInBytes)
	if err != nil {
		log.Println("Failed to add fake data: ", err)
		return "", err
	}

	return cartID, nil
}

func (p *ProductOps) Delete(cartID, productID string) error {
	tx, err := p.DB.Begin()
	if err != nil {
		log.Println("Couldn't start transaction: ", err)
		return err
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.Println("Transcation Rollback failed: ", err1)
			}
		} else {
			if err1 := tx.Commit(); err1 != nil {
				log.Println("Transcation Commit failed: ", err1)
			}
		}
	}()

	var productsInDB []schema.Product
	var dbScan []byte

	if err := tx.QueryRow("SELECT product_details FROM cart WHERE cart_id = $1",
		cartID).Scan(&dbScan); err != nil {
		if err == sql.ErrNoRows {
			log.Println("No such cart is available")
			return err
		}
		log.Println("Failed to get products: ", err)
		return err
	}

	if err := json.Unmarshal(dbScan, &productsInDB); err != nil {
		log.Println("Failed: ", err)
		return err
	}

	// Now creating a new []string which won't have product whose ID is productID
	var products []schema.Product
	var found bool
	for _, prod := range productsInDB {
		if prod.ID != productID {
			products = append(products, prod)
		} else {
			found = true
		}
	}

	if !found {
		log.Println("Product Not Present in the Cart")
		return errors.New("product not present in the cart")
	}

	updatedProductsInBytes, err := json.Marshal(products)
	if err != nil {
		log.Println("Failed: ", err)
		return err
	}

	_, err = tx.Exec("UPDATE cart SET product_details = $1 WHERE cart_id = $2", updatedProductsInBytes, cartID)
	if err != nil {
		log.Println("Delete update failed: ", err)
		return err
	}

	return nil
}
