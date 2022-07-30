package choice

import (
	"database/sql"
	"fmt"
	"shop/db"
	"shop/terminal"
)

func ChoiceOfAction(database *sql.DB) {
	var num int
	connectOut := db.Connection{
		DB: database,
	}
	connectIn := terminal.Connection{
		DB: database,
	}
L:
	for {
		fmt.Println("Enter:")
		fmt.Println("1 for input new seller,")
		fmt.Println("2 for output all sellers,")
		fmt.Println("3 for input new product,")
		fmt.Println("4 for output all products,")
		fmt.Println("5 for input new customer,")
		fmt.Println("6 for output all customers,")
		fmt.Println("7 for input new order,")
		fmt.Println("8 for output all oders,")
		fmt.Println("9 for exit.")
		fmt.Scan(&num)
		switch num {
		case 1:
			err := terminal.Connection.InputSeller(connectIn)
			if err != nil {
				panic(err)
			}
		case 2:
			getSeller, err := db.Connection.GetSellers(connectOut)
			if err != nil {
				panic(err)
			}
			terminal.PrintSellers(database, getSeller)
		case 3:
			err := terminal.Connection.InputProduct(connectIn)
			if err != nil {
				panic(err)
			}
		case 4:
			getProduct, err := db.Connection.GetProducts(connectOut)
			if err != nil {
				panic(err)
			}
			terminal.PrintProducts(database, getProduct)
		case 5:
			err := terminal.Connection.InputCustomer(connectIn)
			if err != nil {
				panic(err)
			}
		case 6:
			getCustomers, err := db.Connection.GetCustomers(connectOut)
			if err != nil {
				panic(err)
			}
			terminal.PrintCustomers(database, getCustomers)
		case 7:
			err := terminal.Connection.InputOrder(connectIn)
			if err != nil {
				panic(err)
			}
		case 8:
			getOrders, err := db.Connection.GetOrders(connectOut)
			if err != nil {
				panic(err)
			}
			terminal.PrintOrders(database, getOrders)
		case 9:
			break L
		}

	}

}
