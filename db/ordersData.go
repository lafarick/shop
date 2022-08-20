package db

import (
	"fmt"
	"shop/models"
)

func (c Connection) CreateOrderData(inputOrderData models.OrderData) error {
	_, err := c.DB.Exec("insert into orders_data(order_id, product_id, quantity, date) values ($1, $2, $3, $4)",
		&inputOrderData.OrderID, &inputOrderData.Quantity, &inputOrderData.ProductID, &inputOrderData.Date)
	if err != nil {
		return err
	}
	return nil
}

func (c Connection) GetOrdersData() ([]models.OrderData, error) {
	rows, err := c.DB.Query("select * from orders_data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ordersData := []models.OrderData{}

	for rows.Next() {
		orderData := models.OrderData{}
		err := rows.Scan(&orderData.OrderID, &orderData.Quantity, &orderData.ProductID, &orderData.Date)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ordersData = append(ordersData, orderData)
	}
	if err != nil {
		return nil, err
	}
	return ordersData, nil
}

func (c Connection) UpdateOrdersData() error {
	_, err := c.DB.Exec("update orders_data set quantity = $1 where order_id = $2", 3, 5)
	if err != nil {
		return err
	}
	return nil
}

func (c Connection) AlterOrdersData() error {
	_, err := c.DB.Exec("alter table orders_data add date date")
	if err != nil {
		return err
	}
	return nil
}
