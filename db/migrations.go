package db

import "database/sql"

func CreateTableSellers(db *sql.DB) error {
	query := `
		create table 
		sellers( 
			id serial primary key, 
			name varchar not null, 
			last_name varchar not null, 
			email varchar not null
		)`
	_, err := db.Exec(query)
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
