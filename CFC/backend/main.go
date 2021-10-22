package main

import (
	"CFC/backend/CFC/backend/DB"
	Facade "CFC/backend/CFC/backend/facade"
	"fmt"
)

func main() {
	d := *DB.NewDatabaseConnection("sql5446146", "WUi5dvp7gj", "sql5.freemysqlhosting.net:3306", "sql5446146")
	pf := *Facade.NewPersonFacade(d)
	p := *pf.GetPerson(5)
	fmt.Printf("The user with id 5 is %s", p.FirstName())
}
