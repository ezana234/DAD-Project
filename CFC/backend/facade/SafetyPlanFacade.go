package facade

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
)

type SafetyPlanFacade struct {
	safetyDao DAO.SafetyPlanDao
	auth      Auth.AuthenticationManager
}

func NewSafetyPlanFacade(db DB.DatabaseConnection) *SafetyPlanFacade {
	return &SafetyPlanFacade{safetyDao: *DAO.NewSafetyPlanDao(db)}
}

func (spf *SafetyPlanFacade) GetSafetyPlan(safetyId int) (*Model.SafetyPlan, error) {
	return spf.safetyDao.GetByID(safetyId), nil
}

func (spf *SafetyPlanFacade) GetSafetyPlanByUserID(userId int) ([]*Model.SafetyPlan, error) {
	return spf.safetyDao.GetByUserID(userId), nil
}
