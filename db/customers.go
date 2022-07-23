package db

import (
	"database/sql"
	"fmt"
	"shop/models"
)

func CreateCustomer(db *sql.DB, inputCustomer models.Customer) error {
	_, err := db.Exec("insert into customers(name, last_name, email) values ( $1, $2, $3)",
		&inputCustomer.Name, &inputCustomer.LastName, &inputCustomer.Email)
	if err != nil {
		return err
	}
	return nil
}

func GetCustomers(db *sql.DB) ([]models.Customer, error) {
	rows, err := db.Query("select * from customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customers := []models.Customer{}

	for rows.Next() {
		customer := models.Customer{}
		err := rows.Scan(&customer.ID, &customer.Name, &customer.LastName, &customer.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		customers = append(customers, customer)
	}
	if err != nil {
		return nil, err
	}
	return customers, nil

}
