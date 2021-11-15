package model

type Appointment struct {
	AppointmentID     int
	AppointmentTime   string
	AppointmentMedium string
	ClientID          int
	ClinicianID       int
}

func NewAppointment(appointmentTime string, appointmentMedium string, clientID int, clinicianID int) *Appointment {
	return &Appointment{AppointmentTime: appointmentTime, AppointmentMedium: appointmentMedium, ClientID: clientID, ClinicianID: clinicianID}
}

func (a *Appointment) GetAppointmentID() int {
	return a.AppointmentID
}

func (a *Appointment) SetAppointmentID(appointmentID int) {
	a.AppointmentID = appointmentID
}

func (a *Appointment) GetAppointmentTime() string {
	return a.AppointmentTime
}

func (a *Appointment) SetAppointmentTime(appointmentTime string) {
	a.AppointmentTime = appointmentTime
}

func (a *Appointment) GetAppointmentMedium() string {
	return a.AppointmentMedium
}

func (a *Appointment) SetAppointmentMedium(appointmentMedium string) {
	a.AppointmentMedium = appointmentMedium
}

func (a *Appointment) GetClientID() int {
	return a.ClientID
}

func (a *Appointment) SetClientID(clientID int) {
	a.ClientID = clientID
}

func (a *Appointment) GetClinicianID() int {
	return a.ClinicianID
}

func (a *Appointment) SetClinicianID(clinicianID int) {
	a.ClinicianID = clinicianID
}
