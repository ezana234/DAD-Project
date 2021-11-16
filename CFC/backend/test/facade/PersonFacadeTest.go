package facade

import (
	"CFC/backend/CFC/backend/DB"
	Facade "CFC/backend/CFC/backend/facade"
	Model "CFC/backend/CFC/backend/model"
	//"testing"
)

var db DB.DatabaseConnection
var pf Facade.PersonFacade
var p Model.Person
var pCreated Model.Person

func CreateForEach(setUp func(), tearDown func()) func(func()) {
	return func(testFunc func()) {
		setUp()
		testFunc()
		tearDown()
	}
}

var RunTest = CreateForEach(setUp, tearDown)

func setUp() {
	//pf = *Facade.NewPersonFacade(db)
	//p = *Model.NewPerson()
	//pCreated = *Model.NewPerson()
}

func tearDown() {

}
