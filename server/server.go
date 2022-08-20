package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"shop/db"
	"shop/models"

	"github.com/gorilla/mux"
)

type Handlers struct {
	Router *mux.Router
	DB     *sql.DB
}

func ServerMain(handlers Handlers) {
	r := mux.NewRouter()
	r.HandleFunc("/sellerin", handlers.SellerIn)
	r.HandleFunc("/sellerout", handlers.SellersOut)
	r.HandleFunc("/productin", handlers.ProductIn)
	r.HandleFunc("/productout", handlers.ProductsOut)
	r.HandleFunc("/customerin", handlers.CustomerIn)
	r.HandleFunc("/customersout", handlers.CustomersOut)
	r.HandleFunc("/orderin", handlers.OrderIn)
	r.HandleFunc("/ordersout", handlers.OrdersOut)
	r.HandleFunc("/orderdatain", handlers.OrderDataIn)
	r.HandleFunc("/ordersdataout", handlers.OrdersDataOut)
	r.HandleFunc("/getbyid/{id}", handlers.OrderByIDOut).Methods("GET")

	http.ListenAndServe(":8080", r)

}

func (h Handlers) SellerIn(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
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

func (h Handlers) SellersOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
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

func (h Handlers) CustomerIn(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
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

func (h Handlers) CustomersOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
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

func (h Handlers) ProductIn(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
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

func (h Handlers) ProductsOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
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

func (h Handlers) OrderIn(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
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

func (h Handlers) OrdersOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
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

func (h Handlers) OrderDataIn(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
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

func (h Handlers) OrdersDataOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
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

func (h Handlers) OrderByIDOut(w http.ResponseWriter, r *http.Request) {
	connect := db.Connection{
		DB: h.DB,
	}

	var ID int
	OrderData, err := db.Connection.GetOrderDataByID(connect, ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	getOrderData := []OrdersData{}
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
