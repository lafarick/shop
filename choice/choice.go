package choice

import (
	"database/sql"
	"fmt"
	"shop/db"
	"shop/terminal"
)

func ChoiceOfAction(database *sql.DB) {
	var num int
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
			err := terminal.InputSeller(database)
			if err != nil {
				panic(err)
			}
		case 2:
			//fmt.Println("Input the numbe:")
			//fmt.Scan(&ID)
			getSeller, err := db.GetSellers(database)
			if err != nil {
				panic(err)
			}
			terminal.PrintSellers(database, getSeller)
		case 3:
			err := terminal.InputProduct(database)
			if err != nil {
				panic(err)
			}
		case 4:
			getProduct, err := db.GetProducts(database)
			if err != nil {
				panic(err)
			}
			terminal.PrintProducts(database, getProduct)
		case 5:
			err := terminal.InputCustomer(database)
			if err != nil {
				panic(err)
			}
		case 6:
			getCustomers, err := db.GetCustomers(database)
			if err != nil {
				panic(err)
			}
			terminal.PrintCustomers(database, getCustomers)
		case 7:
			err := terminal.InputOrder(database)
			if err != nil {
				panic(err)
			}
		case 8:
			getOrders, err := db.GetOrders(database)
			if err != nil {
				panic(err)
			}
			terminal.PrintOrders(database, getOrders)
		case 9:
			break L
		}

	}
}
