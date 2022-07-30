package db

import (
	"database/sql"
	"fmt"
	"shop/models"
)

func (c Connection) CreateOrder(inputOrder models.Order) error {
	_, err := c.DB.Exec("insert into orders(customer_id, product_id, quantity) values ($1, $2, $3)",
		&inputOrder.CustomerID, &inputOrder.ProductID, &inputOrder.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (c Connection) GetOrders() ([]models.Order, error) {
	rows, err := c.DB.Query("select * from orders")
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

func UpdateOrders(db *sql.DB) error {
	_, err := db.Exec("update orders set id = $1 where id = $2", 4, 7)
	if err != nil {
		return err
	}
	return nil
}
