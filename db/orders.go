package db

import (
	"fmt"
	"shop/models"
)

func (c Connection) CreateOrder(inputOrder models.Order) error {

	_, err := c.DB.Exec("insert into orders(customer_id) values ($1)",
		&inputOrder.CustomerID)
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
		err := rows.Scan(&order.ID, &order.CustomerID)
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

func (c Connection) UpdateOrders() error {
	_, err := c.DB.Exec("update orders set id = $1 where id = $2", 6, 9)
	if err != nil {
		return err
	}
	return nil
}

func (c Connection) AlterOrders() error {
	_, err := c.DB.Exec("alter table orders drop column quantity")
	if err != nil {
		return err
	}
	return nil
}
