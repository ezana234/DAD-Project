package facade

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
)

type ClinicianFacade struct {
	clinicianDao DAO.ClinicianDao
	authManager  *Auth.AuthenticationManager
}

func NewClinicianFacade(db DB.DatabaseConnection) *ClinicianFacade {
	return &ClinicianFacade{clinicianDao: *DAO.NewClinicianDao(db)}
}

func (cf *ClinicianFacade) GetClinicianAuthManager() *Auth.AuthenticationManager {
	return cf.authManager
}

func (cf *ClinicianFacade) GetClinicianByID(clinicianID int) (*Model.Clinician, int) {
	c, err := cf.clinicianDao.GetClinicianByID(clinicianID)
	if err != nil {
		return new(Model.Clinician), 0
	}

	return c, 1
}

func (cf *ClinicianFacade) GetAllClinicians() ([]*Model.Clinician, int) {
	var emptyList []*Model.Clinician

	cList, err := cf.clinicianDao.GetAllClinicians()
	if err != nil {
		return emptyList, 0
	}

	return cList, 1
}

func (cf *ClinicianFacade) GetAllClients() ([]*Model.Person, int) {
	var emptyList []*Model.Person

	cList, err := cf.clinicianDao.GetAllClients()
	if err != nil {
		return emptyList, 0
	}

	return cList, 1
}

func (cf *ClinicianFacade) AddClinician(c Model.Clinician) int {
	c.SetClinicianID(cf.clinicianDao.GetNextClinicianID())

	err := cf.clinicianDao.AddClinician(c)
	if err != nil {
		return 0
	}

	return 1
}

func (cf *ClinicianFacade) UpdateClinician(clinicianID int, c *Model.Clinician) int {
	err := cf.clinicianDao.UpdateClinician(clinicianID, c)
	if err != nil {
		return 0
	}

	return 1
}

func (cf *ClinicianFacade) DeleteClinician(clinicianID int) int {
	err := cf.clinicianDao.DeleteClinician(clinicianID)
	if err != nil {
		return 0
	}

	return 1
}

func (cf *ClinicianFacade) GetUserByClinicianID(clinicianID int) (*Model.Person, int) {
	p, err := cf.clinicianDao.GetUserByClinicianID(clinicianID)
	if err != nil {
		return new(Model.Person), 0
	}

	return p, 1
} 

func (cf *ClinicianFacade) GetClinicianByReferral(referral string) (*Model.Clinician, int) {
	c, err := cf.clinicianDao.GetClinicianByReferral(referral)
	if err != nil {
		return new(Model.Clinician), 0
	}

	return c, 1
}

func (cf *ClinicianFacade) GetSafetyPlansByClinicianID(clinicianID int) ([]*Model.SafetyPlan, int) {
	var emptyList []*Model.SafetyPlan

	spList, err := cf.clinicianDao.GetSafetyPlansByClinicianID(clinicianID)
	if err != nil {
		return emptyList, 0
	}

	return spList, 1
}

func (cf *ClinicianFacade) GetClientUsersByClinicianID(clinicianID int) ([]*Model.Person, int) {
	var emptyList []*Model.Person
	pList, err := cf.clinicianDao.GetClientUsersByClinicianID(clinicianID)
	if err != nil {
		return emptyList, 0
	}

	return pList, 1
}
