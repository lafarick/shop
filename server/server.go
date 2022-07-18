package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"shop/db"
	"shop/models"
)

type Handlers struct {
	DB *sql.DB
}

func ServerMain(handlers Handlers) {
	http.HandleFunc("/in", handlers.AddNewSeller)
	http.HandleFunc("/out", handlers.SellersOut)

	http.ListenAndServe(":8080", nil)

}

func (h Handlers) AddNewSeller(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		fmt.Println("Body is nil")
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	seller := CreateSeller{}
	err = json.Unmarshal(b, &seller)
	if err != nil {
		fmt.Println("failed to unmarshal json, ", err)
		return
	}
	InputSeller := models.Seller{
		ID:       seller.ID,
		Name:     seller.Name,
		LastName: seller.LastName,
		Email:    seller.Email,
	}
	db.CreateSeller(h.DB, InputSeller)

}

func (h Handlers) SellersOut(w http.ResponseWriter, r *http.Request) {
	Seller, err := db.GetSellers(h.DB)
	if err != nil {
		fmt.Println(err)
		return
	}

	getSell := []Sellers{}
	for _, g := range Seller {
		getS := Sellers{
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

/*func (h Handlers) SellerIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
}*/

/*func (h Handlers) SellerInput(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var seller Sellers
	//var inputSeller []models.Seller
	json.Unmarshal(body, &seller)

	json.NewEncoder(w).Encode(seller)

	newSeller, err := json.Marshal(seller)

	inputSeller:=models.Seller{
		Name: ,
	}

	if err != nil {
		fmt.Println(err)
	} else {
		db.CreateSeller(h.DB, )
	}
	/*
		fmt.Println("metod:", r.Method)
		if r.Method == "GET" {
			t, _ := template.ParseFiles("login.gtpl")
			t.Execute(w, nil)
		} else {
			r.ParseForm()
			a, err := r.Form["name"]
			b, err := r.Form["lastname"]
			c, err := r.Form["email"]
			if err != nil {
				fmt.Println(err)
			}
			inputSeller := models.Seller{
				Name:     a,
				LastName: b,
				Email:    c,
			}
		}

}*/
