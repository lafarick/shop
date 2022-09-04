package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"shop/db"
	"shop/models"
	"strconv"
	"strings"
)

type Connection struct {
	DB *sql.DB
}

func ServerMain(c Connection) {
	http.HandleFunc("/sellerin", c.SellerIn)
	http.HandleFunc("/sellerout", c.SellersOut)
	http.HandleFunc("/productin", c.ProductIn)
	http.HandleFunc("/productout", c.ProductsOut)
	http.HandleFunc("/customerin", c.CustomerIn)
	http.HandleFunc("/customersout", c.CustomersOut)
	http.HandleFunc("/orderin", c.OrderIn)
	http.HandleFunc("/ordersout", c.OrdersOut)
	http.HandleFunc("/orderdatain", c.OrderDataIn)
	http.HandleFunc("/ordersdataout", c.OrdersDataOut)
	http.HandleFunc("/getbyid/{id}", c.OrderByIDOut)

	http.ListenAndServe(":8080", nil)

}

func (c Connection) SellerIn(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	seller := CreateSeller{}
	err = json.Unmarshal(body, &seller)
	if err != nil {
		fmt.Println(err)
		return
	}

	InputSeller := models.Seller{
		Name:     seller.Name,
		LastName: seller.LastName,
		Email:    seller.Email,
	}

	db.Connection.CreateSeller(connect, InputSeller)

}

func (c Connection) SellersOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	Seller, err := db.Connection.GetSellers(connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	getSell := []GetSellers{}
	for _, g := range Seller {
		getS := GetSellers{
			ID:       g.ID,
			Name:     g.Name,
			LastName: g.LastName,
		}
		getSell = append(getSell, getS)
	}

	getS, err := json.Marshal(getSell)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(getS)
}

func (c Connection) CustomerIn(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	customer := CreateCustomer{}
	err = json.Unmarshal(b, &customer)
	if err != nil {
		fmt.Println(err)
		return
	}

	InputCustomer := models.Customer{
		Name:     customer.Name,
		LastName: customer.LastName,
		Email:    customer.Email,
	}

	db.Connection.CreateCustomer(connect, InputCustomer)
}

func (c Connection) CustomersOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	Customer, err := db.Connection.GetCustomers(connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	getCustomers := []GetCustomers{}
	for _, g := range Customer {
		getC := GetCustomers{
			ID:       g.ID,
			Name:     g.Name,
			LastName: g.LastName,
			Email:    g.Email,
		}
		getCustomers = append(getCustomers, getC)
	}

	getC, err := json.Marshal(getCustomers)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(getC)
}

func (c Connection) ProductIn(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	product := CreateProduct{}
	err = json.Unmarshal(b, &product)
	if err != nil {
		fmt.Println(err)
		return
	}

	InputProduct := models.Product{
		Name:     product.Name,
		Price:    product.Price,
		SellerID: product.SellerID,
	}

	db.Connection.CreateProduct(connect, InputProduct)
}

func (c Connection) ProductsOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	Product, err := db.Connection.GetProducts(connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	getProducts := []GetProducts{}
	for _, g := range Product {
		getP := GetProducts{
			ID:       g.ID,
			Name:     g.Name,
			Price:    g.Price,
			SellerID: g.SellerID,
		}
		getProducts = append(getProducts, getP)
	}

	getP, err := json.Marshal(getProducts)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(getP)

}

func (c Connection) OrderIn(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	order := CreateOrder{}
	err = json.Unmarshal(b, &order)
	if err != nil {
		fmt.Println(err)
		return
	}

	InputOrder := models.Order{
		CustomerID: order.CustomerID,
	}

	db.Connection.CreateOrder(connect, InputOrder)
}

func (c Connection) OrdersOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	Order, err := db.Connection.GetOrders(connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	getOrders := []GetOrders{}
	for _, g := range Order {
		getO := GetOrders{
			ID:         g.ID,
			CustomerID: g.CustomerID,
		}
		getOrders = append(getOrders, getO)
	}

	getO, err := json.Marshal(getOrders)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(getO)
}

func (c Connection) OrderDataIn(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	orderData := OrdersData{}
	err = json.Unmarshal(b, &orderData)
	if err != nil {
		fmt.Println(err)
		return
	}

	InputOrderData := models.OrderData{
		OrderID:   orderData.OrderID,
		Quantity:  orderData.Quantity,
		ProductID: orderData.ProductID,
		Date:      orderData.Date,
	}

	db.Connection.CreateOrderData(connect, InputOrderData)
}

func (c Connection) OrdersDataOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	OrderData, err := db.Connection.GetOrdersData(connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	getOrdersData := []OrdersData{}
	for _, g := range OrderData {
		getOD := OrdersData{
			OrderID:   g.OrderID,
			Quantity:  g.Quantity,
			ProductID: g.ProductID,
			Date:      g.Date,
		}
		getOrdersData = append(getOrdersData, getOD)
	}

	getOD, err := json.Marshal(getOrdersData)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(getOD)
}

func (c Connection) OrderByIDOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: c.DB,
	}
	var ID, defaultCode int
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println(defaultCode, p[0])
	} else if len(p) > 1 {
		code, err := strconv.Atoi(p[0])
		if err == nil {
			ID, err = strconv.Atoi(p[1])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(code, p[1])
		} else {
			fmt.Println(defaultCode, p[1])
		}
	} else {
		fmt.Println(defaultCode, "")
	}

	getOrderData := []OrdersData{}
	OrderData, err := db.Connection.GetOrderDataByID(connect, ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, g := range OrderData {
		getOD := OrdersData{
			OrderID:   g.OrderID,
			Quantity:  g.Quantity,
			ProductID: g.ProductID,
			Date:      g.Date,
		}
		getOrderData = append(getOrderData, getOD)
	}

	getOD, err := json.Marshal(getOrderData)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(getOD)
}
