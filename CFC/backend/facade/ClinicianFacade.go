package facade

import (
	"CFC/backend/CFC/backend/DB"
	DAO "CFC/backend/CFC/backend/dao"
	//Model "CFC/backend/CFC/backend/model"
)

type ClinicianFacade struct {
	clinicianDao DAO.ClinicianDao
}

func NewClinicianFacade(db DB.DatabaseConnection) *ClinicianFacade {
	return &ClinicianFacade{clinicianDao: *DAO.NewClinicianDao(db)}
}
