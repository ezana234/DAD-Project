package facadeTest

import (
	"CFC/backend/CFC/backend/DB"
	Facade "CFC/backend/CFC/backend/facade"
	Model "CFC/backend/CFC/backend/model"
	"testing"
)

// func CreateForEach(setUp func(), tearDown func()) func(func()) {
// 	return func(testFunc func()) {
// 		setUp()
// 		testFunc()
// 		tearDown()
// 	}
// }

var newUserID int

func setUpPersonFacadeTests() *Facade.PersonFacade {
	var db = DB.NewDatabaseConnection("ydmscaoenbipqz", "f9ac329ae1c957bdd5015e4f91bb7968850dd6eb2773105ff6f2b4efb036de47", "ec2-52-54-237-144.compute-1.amazonaws.com", "5432", "d85fspl6bklvdv")
	pf := Facade.NewPersonFacade(*db)
	return pf
}

func TestAddPersonSuccess(t *testing.T) {
	pNew := *Model.NewPerson("tuser", "tpassword", "Test", "User", "tuserDELETEME@gmail.com", "123 Street", "123456789", "1", "11/25/2021", "04/03/2002")

	pf := setUpPersonFacadeTests()
	intReturn := pf.AddPerson(pNew)
	if intReturn == 0 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}

	newUserID = intReturn

	return
}

func TestGetPersonSuccess(t *testing.T) {
	pf := setUpPersonFacadeTests()
	// print("Running TestGetUserByID...")
	pReturn, intReturn := pf.GetPerson(newUserID)
	if intReturn == 0 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}

	if pReturn.GetUserID() != newUserID {
		t.Errorf("Return userID = %d; want %d", pReturn.GetUserID(), newUserID)
	}
	return
}

func TestGetAllSuccess(t *testing.T) {
	pf := setUpPersonFacadeTests()
	// print("Running TestGetAll... ")
	pListReturn, intReturn := pf.GetAllPersons()
	if intReturn != 1 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}

	if len(pListReturn) <= 0 {
		t.Errorf("Return pList is empty")
	}

	return
}

//func TestGetPersonByEmailSuccess(t *testing.T) {
//	var pExpectedEmail = "tuserDELETEME@gmail.com"
//
//	pf := setUpPersonFacadeTests()
//	pReturn, intReturn := pf.LoginPersonByEmail(pExpectedEmail, "tpassword")
//	if intReturn != 1 {
//		t.Errorf("Return int = %d; want 1", intReturn)
//	}
//
//	if pReturn.GetUserID() != newUserID {
//		t.Errorf("Return userID = %d; want %d", pReturn.GetUserID(), newUserID)
//	}
//
//	if pReturn.GetEmail() != pExpectedEmail {
//		t.Errorf("Return email = %s; want %s", pReturn.GetEmail(), pExpectedEmail)
//	}
//
//	return
//}

func TestDeletePersonSuccess(t *testing.T) {
	pf := setUpPersonFacadeTests()

	intReturn := pf.DeletePerson(newUserID)
	if intReturn == 0 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}

	return
}

func TestGetClientNameByUserID(t *testing.T) {
	pf := setUpPersonFacadeTests()

	clientName, success := pf.GetClientNameByUserID(2)
	if success == 0 {
		t.Errorf("Return int = %d; want 1", success)
	}
	println(clientName.Print())
	if clientName.GetFirstName() != "Tessi" {
		t.Errorf("Return client first name = \"%s\"; want \"%s\"", clientName.GetFirstName(), "Tessi")
	}

}
