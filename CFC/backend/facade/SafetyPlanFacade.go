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

func NewSafetyPlanFacade(db DB.DBConnection) *SafetyPlanFacade {
	return &SafetyPlanFacade{safetyDao: *DAO.NewSafetyPlanDao(db)}
}

func (spf *SafetyPlanFacade) GetSafetyPlan(safetyId int) (*Model.SafetyPlan, error) {
	return spf.safetyDao.GetByID(safetyId), nil
}

func (spf *SafetyPlanFacade) GetSafetyPlanByUserID(userId int) ([]*Model.SafetyPlan, error) {
	return spf.safetyDao.GetByUserID(userId), nil
}

func (spf *SafetyPlanFacade) GetAllSafetyPlans() ([]*Model.SafetyPlan, int) {
	var emptyList []*Model.SafetyPlan

	spList, err := spf.safetyDao.GetAll()
	if err != nil {
		return emptyList, 0
	}

	return spList, 1
}

func (spf *SafetyPlanFacade) AddSafetyPlan(sp *Model.SafetyPlan) int {
	sp.SetSafetyID(spf.safetyDao.GetNextSafetyID())
	rowsAffected, err := spf.safetyDao.Add(sp)
	if err != nil {
		return 0
	}
	if rowsAffected <= 0 {
		return -1
	}

	return 1
}

func (spf *SafetyPlanFacade) UpdateSafetyPlan(safetyID int, sp *Model.SafetyPlan) int {
	rowsAffected, err := spf.safetyDao.Update(safetyID, sp)
	if err != nil || rowsAffected <= 0 {
		return 0
	}

	return 1
}

func (spf *SafetyPlanFacade) DeleteSafetyPlan(safetyID int) int {
	rowsAffected, err := spf.safetyDao.Delete(safetyID)
	if err != nil {
		return 0
	}
	if rowsAffected <= 0 {
		return -1
	}

	return 1
}
