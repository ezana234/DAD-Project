package facade

import (
	"CFC/backend/CFC/backend/DB"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
)

type ClinicianFacade struct {
	clinicianDao DAO.ClinicianDao
}

func NewClinicianFacade(db DB.DatabaseConnection) *ClinicianFacade {
	return &ClinicianFacade{clinicianDao: *DAO.NewClinicianDao(db)}
}

func (cf *ClinicianFacade) GetClinician(clinicianID int) *Model.Clinician {
	return cf.clinicianDao.GetClinician(clinicianID)
}

func (cf *ClinicianFacade) GetClinicians() []*Model.Clinician {
	return cf.clinicianDao.GetAllClinicians()
}

func (cf *ClinicianFacade) AddClinician(c Model.Clinician) error {
	c.SetClinicianID(cf.clinicianDao.GetNextID())
	_ = cf.clinicianDao.AddClinician(c)
	return nil
}

func (cf *ClinicianFacade) UpdateClinician(clinicianID int, c *Model.Clinician) error {
	_ = cf.clinicianDao.UpdateClinician(clinicianID, c)
	return nil
}

func (cf *ClinicianFacade) DeleteClinician(clinicianID int) error {
	_ = cf.clinicianDao.DeleteClinician(clinicianID)
	return nil
}
