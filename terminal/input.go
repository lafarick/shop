package terminal

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"shop/db"
	"shop/models"
)

func InputSeller(DB *sql.DB) error {
	var seller models.Seller
	fmt.Println("Input data:")

	_, err := fmt.Scan(&seller.Name, &seller.LastName, &seller.Email)
	if err != nil {
		return err
	}

	err = db.CreateSeller(DB, seller)
	if err != nil {
		return err
	}

	return nil
}

func InputProduct(DB *sql.DB) error {
	var product models.Product
	cin := bufio.NewScanner(os.Stdin)
	fmt.Println("Input name:")
	cin.Scan()
	product.Name = cin.Text()
	if err := cin.Err(); err != nil {

		return err
	}

	fmt.Println("Input price:")
	fmt.Scanln(&product.Price)
	fmt.Println("Input seller ID:")
	fmt.Scanln(&product.SellerID)

	err := db.CreateProduct(DB, product)
	if err != nil {
		return err
	}

	return nil
}
