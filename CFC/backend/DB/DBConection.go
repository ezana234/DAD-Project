package DB

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//func init() {
//	sql.Register("mysql", &MySQLDriver{})
//}

type DatabaseConnection struct {
	username string
	password string
	hostname string
	dbname   string
	dsn      string
}

func NewDatabaseConnection(username string, password string, hostname string, dbname string) *DatabaseConnection {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	return &DatabaseConnection{
		username: username,
		password: password,
		hostname: hostname,
		dbname:   dbname,
		dsn:      dsn,
	}
}

func (d *DatabaseConnection) Select(query *NamedParameterQuery, parameterMap map[string]interface{}) ([][]string, error) {
	var rs [][]string

	query.SetValuesFromMap(parameterMap)
	db, err := sql.Open("mysql", d.dsn)
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return rs, errors.New("failed to connect to DB")
	}

	rows, err := db.Query(query.GetParsedQuery(), query.GetParsedParameters()...)
	if err != nil {
		log.Printf("Error %s when querying DB\n", err)
		return rs, errors.New("failed to query database")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error %s when closing DB\n", err)
		}
	}(db)

	cols, err := rows.Columns()
	if err != nil {
		log.Printf("Error %s when calculating column count", err)
		return rs, errors.New("failed to get columns")
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

func (d *DatabaseConnection) Update(query *NamedParameterQuery, parameterMap map[string]interface{}) error {
	query.SetValuesFromMap(parameterMap)
	db, err := sql.Open("mysql", d.dsn)
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return errors.New("failed to connect to DB")
	}

	_, err = db.Exec(query.GetParsedQuery(), query.GetParsedParameters()...)
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
