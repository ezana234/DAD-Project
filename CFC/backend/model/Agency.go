package model

type Agency struct {
	AgencyId       int
	Name           string
	PhoneNumber    string
	Specialization string
}

func NewAgency(name string, phoneNumber string, specialization string) *Agency {
	return &Agency{Name: name, PhoneNumber: phoneNumber, Specialization: specialization}
}
