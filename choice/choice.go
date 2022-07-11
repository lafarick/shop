package choice

import (
	"database/sql"
	"fmt"
	"shop/db"
	"shop/serverHello"
	"shop/serverInc"
	"shop/terminal"
)

func ChoiceOfAction(database *sql.DB) {
	var num, ID int
L:
	for {
		fmt.Println("Enter:")
		fmt.Println("1 for input new seller,")
		fmt.Println("2 for output all sellers,")
		fmt.Println("3 for input new product,")
		fmt.Println("4 for output specific seller`s products.")
		fmt.Println("5 for start Hello server;")
		fmt.Println("6 for start Increment server;")
		fmt.Println("7 for exit.")
		fmt.Scan(&num)
		switch num {
		case 1:
			err := terminal.InputSeller(database)
			if err != nil {
				panic(err)
			}
		case 2:
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
			fmt.Println("Input the numbe:")
			fmt.Scan(&ID)
			getProduct, err := db.GetProducts(database, ID)
			if err != nil {
				panic(err)
			}
			terminal.PrintProducts(database, getProduct)
		case 5:
			serverHello.ConnectHello()
		case 6:
			serverInc.ConnectInc()
		case 7:
			break L
		}

	}
}
