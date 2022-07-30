package db

import (
	"fmt"
	"shop/models"
)

func (c Connection) CreateSeller(inputSeller models.Seller) error {
	_, err := c.DB.Exec("insert into sellers(name,last_name, email) values ( $1, $2, $3)",
		&inputSeller.Name, &inputSeller.LastName, &inputSeller.Email)
	if err != nil {
		return err
	}
	return err
}

func (c Connection) GetSellers() ([]models.Seller, error) {
	rows, err := c.DB.Query("select * from sellers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sellers := []models.Seller{}

	for rows.Next() {
		seller := models.Seller{}
		err := rows.Scan(&seller.ID, &seller.Name, &seller.LastName, &seller.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sellers = append(sellers, seller)
	}
	if err != nil {
		return nil, err
	}
	return sellers, nil
}
