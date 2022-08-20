package db

import (
	"fmt"
	"shop/models"
)

func (c Connection) GetOrderDataByID(orderID int) ([]models.OrderData, error) {
	rows, err := c.DB.Query("select order_id, product_id, quantity, date from orders_data where order_id = $1", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersData := []models.OrderData{}

	for rows.Next() {
		oD := models.OrderData{}
		err := rows.Scan(&oD.OrderID, &oD.ProductID, &oD.Quantity, &oD.Date)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ordersData = append(ordersData, oD)
	}
	return ordersData, nil
}
