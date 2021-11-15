package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
	"strconv"
)

// TODO GetClientByAppointmentID()

// TODO GetClinicianByAppointmentID()

type AppointmentDao struct {
	db DB.DatabaseConnection
}

func NewAppointmentDAO(db DB.DatabaseConnection) *AppointmentDao {
	return &AppointmentDao{db: db}
}

func (ad *AppointmentDao) GetAppointmentByID(appID int) (*Model.Appointment, error) {
	var query = "SELECT * FROM cfc.appointments WHERE appointments.appointmentid=$1 LIMIT 1"
	var parameters = []interface{}{appID}

	result, err := ad.db.Select(query, parameters)
	if err != nil {
		return new(Model.Appointment), err
	}

	var res = result[0]
	appid, _ := strconv.ParseInt(res[0], 10, 64)
	clientid, _ := strconv.ParseInt(res[3], 10, 64)
	clinicianid, _ := strconv.ParseInt(res[4], 10, 64)
	app := Model.NewAppointment(res[1], res[2], int(clientid), int(clinicianid))
	app.SetAppointmentID(int(appid))

	return app, nil
}

func (ad *AppointmentDao) AddAppointment(app Model.Appointment) error {
	var query = "INSERT INTO cfc.appointments(appointmentid, appointmenttime, appointmentmedium, client_clientid, clinician_clinicianid) VALUES($1,$2,$3,$4,$5)"
	var parameters = []interface{}{app.GetAppointmentID(), app.GetAppointmentTime(), app.GetAppointmentMedium(), app.GetClientID(), app.GetClinicianID()}
	
	return ad.db.Insert(query, parameters)
}

func (ad *AppointmentDao) UpdateAppointment(appID int, app Model.Appointment) error {
	var query = "UPDATE cfc.appointments SET appointmenttime=$1, appointmentmedium=$2, client_clientid=$3, clinician_clinicianid=$4 WHERE appointmentid=$5"
	var parameters = []interface{}{app.GetAppointmentTime(), app.GetAppointmentMedium(), app.GetClientID(), app.GetClinicianID(), appID}
	
	return ad.db.Update(query, parameters)
}

func (ad *AppointmentDao) DeleteAppointment(appID int) error {
	var query = "DELETE FROM cfc.appointments WHERE appointmentid=$1"
	var parameters = []interface{}{appID}

	return ad.db.Delete(query, parameters)
}

func (ad *AppointmentDao) GetAllAppointments() ([]*Model.Appointment, error) {
	var query = "SELECT * FROM cfc.appointments"
	var parameters = []interface{}{}
	var aList []*Model.Appointment

	result, err := ad.db.Select(query, parameters)
	if err != nil {
		return aList, err
	}

	for _, res := range result {
		appid, _ := strconv.ParseInt(res[0], 10, 64)
		clientid, _ := strconv.ParseInt(res[3], 10, 64)
		clinicianid, _ := strconv.ParseInt(res[4], 10, 64)
		app := Model.NewAppointment(res[1], res[2], int(clientid), int(clinicianid))
		app.SetAppointmentID(int(appid))
		aList = append(aList, app)
	}

	return aList, nil
}