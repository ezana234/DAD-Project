package main

import (
	"CFC/backend/CFC/backend/DB"
	Facade "CFC/backend/CFC/backend/facade"
	Model "CFC/backend/CFC/backend/model"
)

func main() {
	db := *DB.NewDatabaseConnection("sql5446146", "WUi5dvp7gj", "sql5.freemysqlhosting.net:3306", "sql5446146")
	cf := *Facade.NewClinicianFacade(db)
	newClinician := *Model.NewClinician(1002)
	cf.AddClinician(newClinician)

}
