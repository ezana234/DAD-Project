package facade

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
	"errors"
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

func (cf *ClinicianFacade) GetClinician(clinicianID int) (*Model.Clinician, error) {
	if cf.authManager.IsCurrentUserAdmin() || cf.authManager.IsCurrentUserClinician() {
		c, _ := cf.clinicianDao.GetClinicianByID(clinicianID)
		return c, nil
	}

	return new(Model.Clinician), errors.New("unable to get clinician: user has incorrect permissions")
}

func (cf *ClinicianFacade) GetClinicians() ([]*Model.Clinician, error) {
	if cf.authManager.IsCurrentUserAdmin() || cf.authManager.IsCurrentUserClinician() {
		return cf.clinicianDao.GetAll(), nil
	}

	return []*Model.Clinician{}, errors.New("unable to get clinicians: user has incorrect permissions")
}

func (cf *ClinicianFacade) GetAllClients() []*Model.Person {
	return cf.clinicianDao.GetAllClients()
}

func (cf *ClinicianFacade) AddClinician(c Model.Clinician) error {
	if cf.authManager.IsCurrentUserAdmin() || cf.authManager.IsCurrentUserClinician() {
		c.SetClinicianID(cf.clinicianDao.GetNextClinicianID())

		return cf.clinicianDao.AddClinician(c)
	}

	return errors.New("unable to add clinician: user has incorrect permissions")
}

func (cf *ClinicianFacade) UpdateClinician(clinicianID int, c *Model.Clinician) error {
	if cf.authManager.IsCurrentUserAdmin() || cf.authManager.IsCurrentUserClinician() {
		return cf.clinicianDao.UpdateClinician(clinicianID, c)
	}

	return errors.New("unable to update clinician: user has incorrect permissions")
}

func (cf *ClinicianFacade) DeleteClinician(clinicianID int) error {
	if cf.authManager.IsCurrentUserAdmin() {
		return cf.clinicianDao.DeleteClinician(clinicianID)
	}

	return errors.New("unable to delete clinician: user has incorrect permissions")
}

func (cf *ClinicianFacade) GetUserByClinicianID(clinicianID int) (*Model.Person, error) {
	return cf.clinicianDao.GetUserByClinicianID(clinicianID)
} 

func (cf *ClinicianFacade) GetClinicianByReferral(referral string) (*Model.Clinician, error) {
	return cf.clinicianDao.GetClinicianByReferral(referral)
}

func (cf *ClinicianFacade) GetSafetyPlansByClinicianID(clinicianID int) ([]*Model.SafetyPlan, error) {
	return cf.clinicianDao.GetSafetyPlansByClinicianID(clinicianID)
}
