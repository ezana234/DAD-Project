package facadeTest

import (
	"CFC/backend/CFC/backend/DB"
	"CFC/backend/CFC/backend/facade"
	Model "CFC/backend/CFC/backend/model"
	"testing"
)

func setUpAppointmentFacadeTests() *facade.AppointmentFacade {
	var db = DB.NewDatabaseConnection("ydmscaoenbipqz", "f9ac329ae1c957bdd5015e4f91bb7968850dd6eb2773105ff6f2b4efb036de47", "ec2-52-54-237-144.compute-1.amazonaws.com", "5432", "d85fspl6bklvdv")
	af := facade.NewAppointmentFacade(*db)

	return af
}

var newAppID int

func TestGetAppointmentByIDSuccess(t *testing.T) {
	af := setUpAppointmentFacadeTests()

	var aExpected = *Model.NewAppointment("", "Zoom", 334, 341)
	aExpected.SetAppointmentID(1)
	aReturn, intReturn := af.GetAppointmentByID(1)
	if intReturn == 0 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}

	if aReturn.GetAppointmentID() != aExpected.GetAppointmentID() {
		t.Errorf("Return appointmentID = %d; want %d", aReturn.GetAppointmentID(), aExpected.GetAppointmentID())
	}

	if aReturn.GetClientID() != aExpected.GetClientID() {
		t.Errorf("Return clientID = %d; want %d", aReturn.GetClientID(), aExpected.GetClientID())
	}

	return
}

func TestAddAppointmentSuccess(t *testing.T) {
	af := setUpAppointmentFacadeTests()
	aNew := *Model.NewAppointment("2020-11-25 11:00:00.000000", "phone", 335, 342)
	intReturn := af.AddAppointment(aNew)
	if intReturn == 0 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}

	newAppID = intReturn

	return
}

func TestDeleteAppointmentSuccess(t *testing.T) {
	af := setUpAppointmentFacadeTests()
	intReturn := af.DeleteAppointment(newAppID)
	if intReturn == 0 {
		t.Errorf("Return int = %d; want 1", intReturn)
	}
	return
}
