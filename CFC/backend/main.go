package main

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	Facade "CFC/backend/CFC/backend/facade"
)

func main() {
	db := *DB.NewDatabaseConnection("root", "password", "127.0.0.1:3306", "sql5446146")
	//db := *DB.NewDatabaseConnection("sql5446146", "WUi5dvp7gj", "sql5.freemysqlhosting.net:3306", "sql5446146")
	auth := *Auth.NewAuthenticationManager()
	pf := *Facade.NewPersonFacade(db, auth)
	pf.LoginPersonByUserName("admin", "password")
	plist, _ := pf.GetPersons()
	for _, a := range plist {
		println(a.GetUserName())
	}

}
