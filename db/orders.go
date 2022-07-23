package db

import (
	"database/sql"
	"fmt"
	"shop/models"
)

func CreateOrder(db *sql.DB, inputOrder models.Order) error {
	_, err := db.Exec("insert into orders(customer_id, product_id, quantity) values ($1, $2, $3)",
		&inputOrder.CustomerID, &inputOrder.ProductID, &inputOrder.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func GetOrders(db *sql.DB) ([]models.Order, error) {
	rows, err := db.Query("select * from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []models.Order{}

	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.ID, &order.CustomerID, &order.ProductID, &order.Quantity)
		if err != nil {
			fmt.Println(err)
			continue
		}
		orders = append(orders, order)
	}
	if err != nil {
		return nil, err
	}
	return orders, nil
}
