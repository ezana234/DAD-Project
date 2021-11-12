package model

type Clinician struct {
	ClinicianID int
	UserID      int
	Referral    string
}

func NewClinician(userID int, referral string) *Clinician {
	return &Clinician{UserID: userID, Referral: referral}
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

func (c *Clinician) GetReferral() string {
	return c.Referral
}

func (c *Clinician) SetReferral(referral string) {
	c.Referral = referral
}
