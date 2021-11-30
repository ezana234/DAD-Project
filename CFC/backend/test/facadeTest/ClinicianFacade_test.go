package facadeTest

import (
	"CFC/backend/CFC/backend/DB"
	Facade "CFC/backend/CFC/backend/facade"
	Model "CFC/backend/CFC/backend/model"
	"testing"
)

func setUpClinicianFacade() *Facade.ClinicianFacade {
	var db = *DB.NewDatabaseConnection("ydmscaoenbipqz", "f9ac329ae1c957bdd5015e4f91bb7968850dd6eb2773105ff6f2b4efb036de47", "ec2-52-54-237-144.compute-1.amazonaws.com", "5432", "d85fspl6bklvdv")
	cf := Facade.NewClinicianFacade(db)

	return cf
}

func TestGetClinicianNameByClinicianIDSuccess(t *testing.T) {
	var expectedClinicianID = 4
	var expectedClinicianUserID = 15
	var expectedClinicianFirstName = "Kira"
	var expectedClinicianLastName = "Taffrey"
	expectedClinicianName := Model.NewPerson("", "", expectedClinicianFirstName, expectedClinicianLastName, "", "", "", "", "", "")
	expectedClinicianName.SetUserID(expectedClinicianUserID)

	cf := setUpClinicianFacade()
	returnClinicianName, success := cf.GetClinicianNameByClinicianID(expectedClinicianID)
	if success == 0 {
		t.Errorf("Return int = %d; want 1", success)
	}
	if returnClinicianName.GetUserID() != 15 {
		t.Errorf("Return userID = %d; want %d", returnClinicianName.GetUserID(), expectedClinicianUserID)
	}
	if returnClinicianName.GetFirstName() != expectedClinicianFirstName {
		t.Errorf("Return firstname = %s; want %s", returnClinicianName.GetFirstName(), expectedClinicianFirstName)
	}
	if returnClinicianName.GetLastName() != expectedClinicianLastName {
		t.Errorf("Return lastname = %s; want %s", returnClinicianName.GetLastName(), expectedClinicianLastName)
	}
	return
}
