package db

import (
	"fmt"
	"shop/models"
)

func (c Connection) CreateProduct(inputProduct models.Product) error {

	_, err := c.DB.Exec("insert into products(name, price, seller_id) values ( $1, $2, $3)",
		&inputProduct.Name, &inputProduct.Price, &inputProduct.SellerID)
	if err != nil {
		return err
	}
	return nil
}

func (c Connection) GetProducts() ([]models.Product, error) {
	rows, err := c.DB.Query("select* from products")
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

/*func UpdateProducts(db *sql.DB) error {
	_, err := db.Exec("update products set price = $1 where id = $2", 1300, 5)
	if err != nil {
		return err
	}
	return nil
}*/
