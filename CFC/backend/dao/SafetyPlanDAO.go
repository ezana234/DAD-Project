package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
	"strconv"
)

type SafetyPlanDao struct {
	db DB.DatabaseConnection
}

func NewSafetyPlanDao(db DB.DatabaseConnection) *SafetyPlanDao {
	return &SafetyPlanDao{db: db}
}

func (spd *SafetyPlanDao) GetByID(safetyID int) *Model.SafetyPlan {
	query := DB.NewNamedParameterQuery("SELECT * FROM safety_plan WHERE safetyId=:safetyID")
	var parameterMap = map[string]interface{}{
		"safetyID": safetyID,
	}

	result, err := spd.db.Select(query, parameterMap)
	if err != nil {
		return new(Model.SafetyPlan)
	}

	spuid, _ := strconv.ParseInt(result[0][0], 10, 64)
	uc, _ := strconv.ParseInt(result[0][6], 10, 64)
	clientuid, _ := strconv.ParseInt(result[0][7], 10, 64)
	clinicianid, _ := strconv.ParseInt(result[0][8], 10, 64)
	sp := Model.NewSafetyPlan(result[0][1], result[0][2], result[0][3], result[0][4], result[0][5], int(uc), int(clientuid), int(clinicianid))
	sp.SetSafetyID(int(spuid))

	return sp
}

func (spd *SafetyPlanDao) GetAll() []*Model.SafetyPlan {
	query := DB.NewNamedParameterQuery("SELECT * FROM safety_plan")
	var spList []*Model.SafetyPlan

	result, err := spd.db.Select(query, map[string]interface{}{})
	if err != nil || len(result) == 0 {
		return spList
	}

	for _, res := range result {
		spuid, _ := strconv.ParseInt(res[0], 10, 64)
		uc, _ := strconv.ParseInt(res[6], 10, 64)
		clientuid, _ := strconv.ParseInt(res[7], 10, 64)
		clinicianid, _ := strconv.ParseInt(res[8], 10, 64)
		sp := Model.NewSafetyPlan(res[1], res[2], res[3], res[4], res[5], int(uc), int(clientuid), int(clinicianid))
		sp.SetSafetyID(int(spuid))
		spList = append(spList, sp)
	}

	return spList
}

func (spd *SafetyPlanDao) Add(sp Model.SafetyPlan) error {
	query := DB.NewNamedParameterQuery("INSERT INTO safety_plan(safetyId,triggers,warningSigns,destructiveBehaviors,internalStrategies,updatedDatetime,updatedClinician,Client_clientId,Clinician_clinicianId) VALUES(:safetyId,:triggers,:warningSigns,:destructiveBehaviors,:internalStrategies,:updatedDatetime,:updatedClinician,:Client_clientId,:Clinician_clinicianId)")
	var parameterMap = map[string]interface{}{
		"safetyId":              sp.GetSafetyID(),
		"triggers":              sp.GetTriggers(),
		"warningSigns":          sp.GetWarningSigns(),
		"destructiveBehaviors":  sp.GetDestructiveBehaviors(),
		"internalStrategies":    sp.GetInternalStrategies(),
		"updatedDatetime":       sp.GetUpdatedDatetime(),
		"updatedClinician":      sp.GetUpdatedClinician(),
		"Client_clientId":       sp.GetClientID(),
		"Clinician_clinicianId": sp.GetClinicianID(),
	}

	return spd.db.Update(query, parameterMap)
}

func (spd *SafetyPlanDao) Update(userID int, sp *Model.SafetyPlan) error {
	query := DB.NewNamedParameterQuery("UPDATE safety_plan SET triggers=:triggers,warningSigns=:warningSigns,destructiveBehaviors=:destructiveBehaviors,internalStrategies=:internalStrategies,updatedDatetime=:updatedDateTime,updatedClinician=:updatedClinician,Client_clientId=:Client_ClientId,Clinician_clinicianId=:Clinician_clinicianId WHERE safetyId=:safetyId")
	var parameterMap = map[string]interface{}{
		"triggers":              sp.GetTriggers(),
		"warningSigns":          sp.GetWarningSigns(),
		"destructiveBehaviors":  sp.GetDestructiveBehaviors(),
		"internalStrategies":    sp.GetInternalStrategies(),
		"updatedDatetime":       sp.GetUpdatedDatetime(),
		"updatedClinician":      sp.GetUpdatedClinician(),
		"Client_clientId":       sp.GetClientID(),
		"Clinician_clinicianId": sp.GetClinicianID(),
		"safetyId":              userID,
	}

	return spd.db.Update(query, parameterMap)
}

func (spd *SafetyPlanDao) Delete(safetyId int) error {
	query := DB.NewNamedParameterQuery("DELETE FROM safety_plan WHERE safetyId=:safetyId")
	var parameterMap = map[string]interface{}{
		"safetyId": safetyId,
	}

	return spd.db.Update(query, parameterMap)
}

// TODO GetClientBySafetyPlanID()

// TODO GetClinicianBySafetyPlanID()
