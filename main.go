package main

import (
	"shop/choice"
	"shop/db"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	/*err = db.CreateTableSellers(database)
	if err != nil {
		panic(err)
	}*/

	/*err = db.CreateTableProducts(database)
	if err != nil {
		panic(err)
	}*/

	choice.ChoiceOfAction(database)
}
