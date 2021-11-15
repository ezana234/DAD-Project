package DB

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	database DatabaseConnection
}

func NewDatabase(connection DatabaseConnection) *Database {
	return &Database{database: connection}
}

type DatabaseConnection struct {
	username string
	password string
	hostname string
	port     string
	dbname   string
	dsn      string
	psqlInfo string
}

func NewDatabaseConnection(username string, password string, hostname string, port string, dbname string) *DatabaseConnection {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require",
		hostname, port, username, password, dbname)
	return &DatabaseConnection{
		username: username,
		password: password,
		hostname: hostname,
		dbname:   dbname,
		dsn:      dsn,
		psqlInfo: psqlInfo,
	}
}

func (d *DatabaseConnection) Select(query string, parameters []interface{}) ([][]string, error) {
	var rs [][]string

	db, err := sql.Open("postgres", d.psqlInfo)
	if err != nil {
		println("oof")
	}

	rows, err := db.Query(query, parameters...)
	if err != nil {
		log.Printf("Error: %s when running query", err)
	}

	cols, err := rows.Columns()
	if err != nil {
		log.Printf("Error %s when calculating column count", err)
	}

	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))

	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			log.Printf("Error %s when scanning row\n", err)
			return rs, errors.New("failed to scan row")
		}

		var rowlist = make([]string, len(cols))
		for i, raw := range rawResult {
			if raw == nil {
				rowlist[i] = "\\N"
			} else {
				rowlist[i] = string(raw)
			}
		}
		rs = append(rs, rowlist)
	}

	return rs, nil
}

func (d *DatabaseConnection) Insert(query string, parameters []interface{}) error {
	db, err := sql.Open("postgres", d.psqlInfo)
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return errors.New("failed to connect to DB")
	}

	_, err = db.Exec(query, parameters...)
	if err != nil {
		log.Printf("Error %s when inserting into DB\n", err)
		return errors.New("failed to insert into database")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error %s when closing DB\n", err)
		}
	}(db)

	return nil
}

func (d *DatabaseConnection) Update(query string, parameters []interface{}) error {
	db, err := sql.Open("postgres", d.psqlInfo)
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return errors.New("failed to connect to DB")
	}

	_, err = db.Exec(query, parameters...)
	if err != nil {
		log.Printf("Error %s when updating DB\n", err)
		return errors.New("failed to update database")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error %s when closing DB\n", err)
		}
	}(db)

	return nil
}

func (d *DatabaseConnection) Delete(query string, parameters []interface{}) error {
	db, err := sql.Open("postgres", d.psqlInfo)
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return errors.New("failed to connect to DB")
	}

	_, err = db.Exec(query, parameters...)
	if err != nil {
		log.Printf("Error %s when deleting from DB\n", err)
		return errors.New("failed to delete from database")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error %s when closing DB\n", err)
		}
	}(db)

	return nil
}
