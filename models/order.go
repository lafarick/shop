package models

type Order struct {
	ID         int
	CustomerID int
	ProductID  int
	Quantity   int
}
