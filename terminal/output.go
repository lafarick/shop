package terminal

import (
	"database/sql"
	"fmt"
	"shop/models"
)

func PrintSellers(DB *sql.DB, sellers []models.Seller) {

	for _, seller := range sellers {
		fmt.Println(seller.ID, seller.Name, seller.LastName, seller.Email)
	}
}

func PrintProducts(DB *sql.DB, products []models.Product) {

	for _, product := range products {
		fmt.Println(product.ID, product.Name, product.Price, product.SellerID)
	}
}

func PrintCustomers(DB *sql.DB, customers []models.Customer) {
	for _, customer := range customers {
		fmt.Println(customer.ID, customer.Name, customer.LastName, customer.Email)
	}
}

func PrintOrders(DB *sql.DB, orders []models.Order) {
	for _, order := range orders {
		fmt.Println(order.ID, order.CustomerID)
	}
}

func (c Connection) PrintOrdersData(ordersData []models.OrderData) {
	for _, orderData := range ordersData {
		fmt.Println(orderData.OrderID, orderData.Quantity, orderData.ProductID, orderData.Date)
	}
}
