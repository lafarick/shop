package db

import (
	"database/sql"
	"fmt"
	"shop/models"
)

func CreateProduct(db *sql.DB, inputProduct models.Product) error {

	_, err := db.Exec("insert into products(name, price, seller_id) values ( $1, $2, $3)",
		&inputProduct.Name, &inputProduct.Price, &inputProduct.SellerID)
	if err != nil {
		return err
	}
	return err
}

func GetProducts(db *sql.DB, ID int) ([]models.Product, error) {
	rows, err := db.Query("select* from products where seller_id=$1", &ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []models.Product{}

	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.SellerID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, product)
	}
	if err != nil {
		return nil, err
	}
	return products, nil
}
