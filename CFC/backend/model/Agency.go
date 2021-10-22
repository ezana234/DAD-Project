package model

type Agency struct {
	agencyID       int
	name           string
	phoneNumber    string
	specialization string
}

func NewAgency(name string, phoneNumber string, specialization string) *Agency {
	return &Agency{name: name, phoneNumber: phoneNumber, specialization: specialization}
}

func (a *Agency) AgencyID() int {
	return a.agencyID
}

func (a *Agency) SetAgencyID(agencyID int) {
	a.agencyID = agencyID
}

func (a *Agency) Name() string {
	return a.name
}

func (a *Agency) SetName(name string) {
	a.name = name
}

func (a *Agency) PhoneNumber() string {
	return a.phoneNumber
}

func (a *Agency) SetPhoneNumber(phoneNumber string) {
	a.phoneNumber = phoneNumber
}

func (a *Agency) Specialization() string {
	return a.specialization
}

func (a *Agency) SetSpecialization(specialization string) {
	a.specialization = specialization
}
