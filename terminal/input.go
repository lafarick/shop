package terminal

import (
	"database/sql"
	"fmt"
	"shop/db"
	"shop/models"
)

func InputSeller(DB *sql.DB) error {
	var seller models.Seller
	fmt.Println("Input data:")

	_, err := fmt.Scan(&seller.Name, &seller.LastName, &seller.Email)
	if err != nil {
		return err
	}

	err = db.CreateSeller(DB, seller)
	if err != nil {
		return err
	}

	return nil
}

func InputProduct(DB *sql.DB) error {
	var product models.Product
	fmt.Println("Input data:")

	_, err := fmt.Scan(&product.Name, &product.Price, &product.SellerID)
	if err != nil {
		return err
	}

	err = db.CreateProduct(DB, product)
	if err != nil {
		return err
	}

	return nil
}
