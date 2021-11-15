package facade

import (
	"CFC/backend/CFC/backend/DB"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
)

type AppointmentFacade struct {
	appointmentDao DAO.AppointmentDao
}

func NewAppointmentFacade(db DB.DatabaseConnection) *AppointmentFacade {
	return &AppointmentFacade{appointmentDao: *DAO.NewAppointmentDAO(db)}
}

func (af *AppointmentFacade) GetAppointmentByID(appID int) (*Model.Appointment, int) {
	app, err := af.appointmentDao.GetAppointmentByID(appID)
	if err != nil {
		return new(Model.Appointment), 0
	}

	return app, 1
}

func (af *AppointmentFacade) AddAppointment(app Model.Appointment) int {
	err := af.appointmentDao.AddAppointment(app)
	if err != nil {
		return 0
	}

	return 1
}

func (af *AppointmentFacade) UpdateAppointment(appID int, app Model.Appointment) int {
	err := af.appointmentDao.UpdateAppointment(appID, app)
	if err != nil {
		return 0
	}

	return 1
}

func (af *AppointmentFacade) DeleteAppointment(appID int) int {
	err := af.appointmentDao.DeleteAppointment(appID)
	if err != nil {
		return 0
	}

	return 1
}

func (af *AppointmentFacade) GetAllAppointments() ([]*Model.Appointment, int) {
	var emptyList []*Model.Appointment

	aList, err := af.appointmentDao.GetAllAppointments()
	if err != nil {
		return emptyList, 0
	}

	return aList, 1
}