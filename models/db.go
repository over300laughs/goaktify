package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Create an exported global variable to hold the database connection pool.
var DB *sql.DB

func InitDB(host, port, user, password, dbName string) error {

	// setup DB connection
	DBURL := fmt.Sprintf("host=%s port=%s user=%s sslmode=disable password=%s dbname=%s", host, port, user, password, dbName)
	var err error

	DB, err = sql.Open("postgres", DBURL)
	if err != nil {
		fmt.Print("Cannot connect to database")
		return err
	}

	// create campaigns table
	_, err = DB.Exec(`DROP TABLE IF EXISTS campaigns`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
        CREATE TABLE campaigns ( id int PRIMARY KEY, 
                                name varchar(255),
                                description varchar,
                                created_on timestamp default current_timestamp,
                                updated_on timestamp default current_timestamp,
                                is_active BOOL NOT NULL );`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`DROP SEQUENCE IF EXISTS campaigns_id_sequence;`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`CREATE SEQUENCE campaigns_id_sequence START 1 INCREMENT 1;`)
	if err != nil {
		return err
	}

	return nil
}
