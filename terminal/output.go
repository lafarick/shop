package terminal

import (
	"database/sql"
	"fmt"
	"shop/models"
)

func PrintSellers(DB *sql.DB, sellers []models.Seller) {

	for _, seller := range sellers {
		fmt.Println(seller.ID, seller.Name, seller.LastName, seller.Email)
	}
}

func PrintProducts(DB *sql.DB, products []models.Product) {

	for _, product := range products {
		fmt.Println(product.ID, product.Name, product.Price, product.SellerID)
	}
}
