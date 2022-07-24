package db

import "database/sql"

func (c Connection) CreateTableSellers() error {
	query := `
		create table 
		sellers( 
			id serial primary key, 
			name varchar not null, 
			last_name varchar not null, 
			email varchar not null
		)`
	_, err := c.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func CreateTableProducts(db *sql.DB) error {
	query := `
		create table 
		products( 
			id serial primary key, 
			name varchar not null, 
			price int not null, 
			seller_id int references sellers(id)
		)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func CreateTableCustomers(db *sql.DB) error {
	query := `
	create table
	customers(
		id serial primary key,
		name varchar not null,
		last_name varchar not null,
		email varchar unique not null
	)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func CreateTableOrders(db *sql.DB) error {
	query := `
	create table
	orders(
		id serial primary key,
		customer_id int references customers(id),
		product_id int references products(id),
		quantity int
	)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
