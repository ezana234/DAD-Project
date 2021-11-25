package facadeTest

import (
	"CFC/backend/CFC/backend/DB"
	Facade "CFC/backend/CFC/backend/facade"
	"testing"
)

func setUpClientFacadeTests() *Facade.ClientFacade {
	db := DB.NewDatabaseConnection("ydmscaoenbipqz", "f9ac329ae1c957bdd5015e4f91bb7968850dd6eb2773105ff6f2b4efb036de47", "ec2-52-54-237-144.compute-1.amazonaws.com", "5432", "d85fspl6bklvdv")
	cf := Facade.NewClientFacade(*db)
	return cf
}

// func init() {
// 	db := DB.NewDatabaseConnection("ydmscaoenbipqz", "f9ac329ae1c957bdd5015e4f91bb7968850dd6eb2773105ff6f2b4efb036de47", "ec2-52-54-237-144.compute-1.amazonaws.com", "5432", "d85fspl6bklvdv")
// 	cf := Facade.NewClientFacade(*db)
// 	// pf := Facade.NewPersonFacade(*db)
// 	// personModel := Model.NewPerson("tuser2", "tpassword", "Test", "User", "tuser2@gmail.com", "123 Street", "123456789", "1", "11/25/2021", "04/03/2002")
// 	// personModel.SetUserID(1006)

// 	// intReturn := pf.AddPerson(*personModel)
// 	// if intReturn != 1 {
// 	// 	fmt.Println("Error setting up test cases. Must be able to insert person first")
// 	// 	return
// 	// }
// 	clientModel := Model.NewClientBoth(2000, 2000)
// 	cf.AddClient(*clientModel)
// }

// Client Facades
func TestGetClientSuccess(t *testing.T) {
	//setup
	cf := setUpClientFacadeTests()

	//test
	var pExpectedClientID = 336
	var pExpectedUserID = 501
	client, success := cf.GetClientByClientID(pExpectedClientID)
	if success == 0 {
		t.Errorf("Return int = %d; want 1", success)
	}
	if client.GetClientID() != pExpectedClientID {
		t.Errorf("Return clientID = %d; want %d", client.GetClientID(), pExpectedClientID)
	}
	if client.GetUserID() != pExpectedUserID {
		t.Errorf("Return userID = %d; want %d", client.GetUserID(), pExpectedUserID)
	}
}

func TestGetAllClientSuccess(t *testing.T) {
	//setup
	cf := setUpClientFacadeTests()
	//test
	clients, success := cf.GetAllClients()
	if success == 0 {
		t.Errorf("Return int = %d; want 1", success)
	}
	if len(clients) == 0 {
		t.Errorf("Return length of clients = %d; want > 1", len(clients))
	}
}
