package model

type Clinician struct {
	ClinicianID int
	UserID      int
}

func NewClinician(userID int) *Clinician {
	return &Clinician{UserID: userID}
}

func (c *Clinician) GetClinicianID() int {
	return c.ClinicianID
}

func (c *Clinician) SetClinicianID(clinicianID int) {
	c.ClinicianID = clinicianID
}

func (c *Clinician) GetUserID() int {
	return c.UserID
}

func (c *Clinician) SetUserID(userID int) {
	c.UserID = userID
}
