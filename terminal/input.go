package terminal

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"shop/db"
	"shop/models"
)

type Connection struct {
	DB *sql.DB
}

func (c Connection) InputSeller() error {
	connect := db.Connection{
		DB: c.DB,
	}
	var seller models.Seller
	fmt.Println("Input data:")

	_, err := fmt.Scan(&seller.Name, &seller.LastName, &seller.Email)
	if err != nil {
		return err
	}

	err = db.Connection.CreateSeller(connect, seller)
	if err != nil {
		return err
	}

	return nil
}

func (c Connection) InputCustomer() error {
	connect := db.Connection{
		DB: c.DB,
	}
	var customer models.Customer
	fmt.Println("Input data:")

	_, err := fmt.Scan(&customer.Name, &customer.LastName, &customer.Email)
	if err != nil {
		return err
	}

	err = db.Connection.CreateCustomer(connect, customer)
	if err != nil {
		return err
	}
	return nil
}

func (c Connection) InputProduct() error {
	connect := db.Connection{
		DB: c.DB,
	}
	var product models.Product
	cin := bufio.NewScanner(os.Stdin)
	fmt.Println("Input name:")
	cin.Scan()
	product.Name = cin.Text()
	if err := cin.Err(); err != nil {

		return err
	}

	fmt.Println("Input price:")
	fmt.Scanln(&product.Price)
	fmt.Println("Input seller ID:")
	fmt.Scanln(&product.SellerID)

	err := db.Connection.CreateProduct(connect, product)
	if err != nil {
		return err
	}

	return nil
}

func (c Connection) InputOrder() error {
	connect := db.Connection{
		DB: c.DB,
	}
	var order models.Order
	fmt.Println("Input customer ID:")
	_, err := fmt.Scanln(&order.CustomerID)
	if err != nil {
		return err
	}

	err = db.Connection.CreateOrder(connect, order)
	if err != nil {
		return err
	}
	return nil
}

func (c Connection) InputOrderData() error {
	connect := db.Connection{
		DB: c.DB,
	}
	var orderData models.OrderData
	fmt.Println("Input order ID:")
	_, err := fmt.Scanln(&orderData.OrderID)
	if err != nil {
		return err
	}

	fmt.Println("Input product ID:")
	_, err = fmt.Scanln(&orderData.ProductID)
	if err != nil {
		return err
	}

	fmt.Println("Input quantity of product:")
	_, err = fmt.Scanln(&orderData.Quantity)
	if err != nil {
		return err
	}

	fmt.Println("Input date of order:")
	_, err = fmt.Scanln(&orderData.Date)
	if err != nil {
		return err
	}

	err = db.Connection.CreateOrderData(connect, orderData)
	if err != nil {
		return err
	}
	return nil
}
