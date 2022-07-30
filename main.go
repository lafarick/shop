package main

import (
	"shop/choice"
	"shop/db"
	"shop/server"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}
	handlers := server.Handlers{
		DB: database,
	}
	go server.ServerMain(handlers)

	/*err = db.CreateTableSellers(database)
	if err != nil {
		panic(err)
	}*/

	/*err = db.CreateTableProducts(database)
	if err != nil {
		panic(err)
	}*/

	/*err = db.CreateTableCustomers(database)
	if err != nil {
		panic(err)
	}*/

	/*err = db.CreateTableOrders(database)
	if err != nil {
		panic(err)
	}*/

	/*err = db.UpdateProducts(database)
	if err != nil {
		panic(err)
	}*/

	/*err = db.UpdateOrders(database)
	if err != nil {
		panic(err)
	}*/

	choice.ChoiceOfAction(database)

}
