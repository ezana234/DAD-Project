package facade

import (
	Auth "CFC/backend/CFC/backend/auth"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
	"errors"
)

type SafetyPlanFacade struct {
	spd  DAO.SafetyPlanDao
	auth Auth.AuthenticationManager
}

func NewSafetyPlanFacade(spd DAO.SafetyPlanDao, auth Auth.AuthenticationManager) *SafetyPlanFacade {
	return &SafetyPlanFacade{spd: spd, auth: auth}
}

func (spf *SafetyPlanFacade) GetSafetyPlan(safetyId int) (*Model.SafetyPlan, error) {
	if spf.auth.IsCurrentUserAdmin() || spf.auth.IsCurrentUserClinician() {
		return spf.spd.GetByID(safetyId), nil
	}

	return new(Model.SafetyPlan), errors.New("user does not have permission")
}
