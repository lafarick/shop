package terminal

import (
	"database/sql"
	"fmt"
	"shop/db"
	"shop/models"
)

func PrintSellers(DB *sql.DB, sellers []models.Seller) error {
	_, err := db.GetSellers(DB)

	for _, seller := range sellers {
		fmt.Println(seller.ID, seller.Name, seller.LastName, seller.Email)
	}
	if err != nil {
		return err
	}
	return nil
}

func PrintProducts(DB *sql.DB, products []models.Product, ID int) error {
	_, err := db.GetProducts(DB, ID)

	for _, product := range products {
		fmt.Println(product.ID, product.Name, product.Price, product.SellerID)
	}
	if err != nil {
		return err
	}
	return nil
}
