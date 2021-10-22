package model

type Appointment struct {
	appointmentID     int
	appointmentTime   string
	appointmentMedium string
	clientID          int
	clinicianID       int
}

func NewAppointment(appointmentTime string, appointmentMedium string, clientID int, clinicianID int) *Appointment {
	return &Appointment{appointmentTime: appointmentTime, appointmentMedium: appointmentMedium, clientID: clientID, clinicianID: clinicianID}
}

func (a *Appointment) AppointmentID() int {
	return a.appointmentID
}

func (a *Appointment) SetAppointmentID(appointmentID int) {
	a.appointmentID = appointmentID
}

func (a *Appointment) AppointmentTime() string {
	return a.appointmentTime
}

func (a *Appointment) SetAppointmentTime(appointmentTime string) {
	a.appointmentTime = appointmentTime
}

func (a *Appointment) AppointmentMedium() string {
	return a.appointmentMedium
}

func (a *Appointment) SetAppointmentMedium(appointmentMedium string) {
	a.appointmentMedium = appointmentMedium
}

func (a *Appointment) ClientID() int {
	return a.clientID
}

func (a *Appointment) SetClientID(clientID int) {
	a.clientID = clientID
}

func (a *Appointment) ClinicianID() int {
	return a.clinicianID
}

func (a *Appointment) SetClinicianID(clinicianID int) {
	a.clinicianID = clinicianID
}
