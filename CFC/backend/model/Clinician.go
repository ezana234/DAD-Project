package model

type Clinician struct {
	clinicianID int
	userID      int
}

func NewClinician(userID int) *Clinician {
	return &Clinician{userID: userID}
}

func (c *Clinician) ClinicianID() int {
	return c.clinicianID
}

func (c *Clinician) SetClinicianID(clinicianID int) {
	c.clinicianID = clinicianID
}

func (c *Clinician) UserID() int {
	return c.userID
}

func (c *Clinician) SetUserID(userID int) {
	c.userID = userID
}
