package db

import "database/sql"

func CreateTable(db *sql.DB) error {
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
