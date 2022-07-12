package main

import (
	"encoding/json"
	"fmt"
	"shop/choice"
	"shop/db"
	"shop/models"
	"shop/serverHello"
)

func main() {
	go serverHello.ConnectHello()
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	jsonString := `[{"first_word": "Say", "second_word": "Hello", "third_word": "World"}]`

	helloWorld := []models.HelloWorld{}

	err = json.Unmarshal([]byte(jsonString), &helloWorld)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, h := range helloWorld {
		fmt.Println(h.FirstWord, h.SecondWord, h.ThirdWord)

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
