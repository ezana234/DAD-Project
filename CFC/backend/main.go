package main

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	Facade "CFC/backend/CFC/backend/facade"
	Model "CFC/backend/CFC/backend/model"
)

func main() {
	//db := *DB.NewDatabaseConnection("postgres", "password", "localhost","5438" , "d85fspl6bklvdv")
	db := *DB.NewDatabaseConnection("ydmscaoenbipqz", "f9ac329ae1c957bdd5015e4f91bb7968850dd6eb2773105ff6f2b4efb036de47", "ec2-52-54-237-144.compute-1.amazonaws.com", "5432", "d85fspl6bklvdv")
	//db := *DB.NewDatabaseConnection("d85fspl6bklvdv", "WUi5dvp7gj", "sql5.freemysqlhosting.net:3306", "d85fspl6bklvdv")
	//auth := *Auth.NewAuthenticationManager()
	p := Model.NewPerson("admin", "password", "admin", "admin", "admin@gmail.com", "123 Admin Street", "1234567890", "4")

	var auth = *Auth.NewAuthenticationManager()
	pf := Facade.NewPersonFacade(db, &auth)
	err := pf.AddPerson(*p)
	if err != nil {
		return
	}
	//adminPassword := Facade.HashPassword("password")
	//p.SetPassword(adminPassword)
	//pf.

}
