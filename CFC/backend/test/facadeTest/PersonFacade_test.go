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

func setUpPersonFacadeTests() *Facade.PersonFacade {
	db := DB.NewDatabaseConnection("ydmscaoenbipqz", "f9ac329ae1c957bdd5015e4f91bb7968850dd6eb2773105ff6f2b4efb036de47", "ec2-52-54-237-144.compute-1.amazonaws.com", "5432", "d85fspl6bklvdv")
	pf := Facade.NewPersonFacade(*db)

	return pf
}

func TestGetPersonSuccess(t *testing.T) {
	pf := setUpPersonFacadeTests()
	// print("Running TestGetUserByID...")
	var pExpected = *Model.NewPerson("tuser", "tpassword", "Test", "User", "tuser@gmail.com", "123 Street", "123456789", "1", "", "04/03/2002")
	pExpected.SetUserID(1006)

	pReturn, intReturn := pf.GetPerson(1006)
	if intReturn == 0 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}

	if pReturn.GetUserID() != pExpected.GetUserID() {
		t.Errorf("Return userID = %d; want %d", pReturn.GetUserID(), pExpected.GetUserID())
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

func TestGetPersonByEmailSuccess(t *testing.T) {
	var pExpectedUserID = 1006
	var pExpectedEmail = "tuser@gmail.com"

	pf := setUpPersonFacadeTests()
	pReturn, intReturn := pf.LoginPersonByEmail(pExpectedEmail, "tpassword")
	if intReturn != 1 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}

	if pReturn.GetUserID() != pExpectedUserID {
		t.Errorf("Return userID = %d; want %d", pReturn.GetUserID(), pExpectedUserID)
	}

	if pReturn.GetEmail() != pExpectedEmail {
		t.Errorf("Return email = %s; want %s", pReturn.GetEmail(), pExpectedEmail)
	}

}

func TestAddPersonSuccess(t *testing.T) {
	var pExpectedUserName = "tuser2"
	pNew := *Model.NewPerson("tuser2", "tpassword", "Test", "User", "tuser2@gmail.com", "123 Street", "123456789", "1", "11/25/2021", "04/03/2002")

	pf := setUpPersonFacadeTests()
	intReturn := pf.AddPerson(pNew)
	if intReturn != 1 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}

	pReturn, intReturn := pf.GetPersonByUserName(pExpectedUserName)
	if intReturn != 1 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}

	if pReturn.GetUserName() != pExpectedUserName {
		t.Errorf("Return username = %s; want %s", pReturn.GetUserName(), pExpectedUserName)
	}

}

//func TestDeletePersonSuccess(t *testing.T) {
//	pf := setUpPersonFacadeTests()
//
//	intReturn := pf.DeletePerson(1006)
//}
