package server

type CreateSeller struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

type GetSellers struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

type CreateProduct struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	SellerID int    `json:"sellerId"`
}

type GetProducts struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	SellerID int    `json:"sellerId"`
}

type CreateCustomer struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

type GetCustomers struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

type CreateOrder struct {
	CustomerID int `json:"customerId"`
	ProductID  int `json:"productId"`
	Quantity   int `json:"quantity"`
}

type GetOrders struct {
	ID         int `json:"id"`
	CustomerID int `json:"customerId"`
	ProductID  int `json:"productId"`
	Quantity   int `json:"quantity"`
}
