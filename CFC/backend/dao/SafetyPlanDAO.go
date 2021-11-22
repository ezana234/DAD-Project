package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
	"fmt"
	"strconv"
)

type SafetyPlanDao struct {
	db DB.DatabaseConnection
}

func NewSafetyPlanDao(db DB.DatabaseConnection) *SafetyPlanDao {
	return &SafetyPlanDao{db: db}
}

func (spd *SafetyPlanDao) GetByID(safetyID int) *Model.SafetyPlan {
	var query = "SELECT * FROM safety_plan WHERE safetyId=$1"
	var parameters = []interface{}{
		safetyID,
	}

	result, err := spd.db.Select(query, parameters)
	if err != nil {
		return new(Model.SafetyPlan)
	}
	fmt.Println(result)
	spuid, _ := strconv.ParseInt(result[0][0], 10, 64)
	uc, _ := strconv.ParseInt(result[0][6], 10, 64)
	clientuid, _ := strconv.ParseInt(result[0][7], 10, 64)
	clinicianid, _ := strconv.ParseInt(result[0][8], 10, 64)
	sp := Model.NewSafetyPlan(result[0][1], result[0][2], result[0][3], result[0][4], result[0][5], int(uc), int(clientuid), int(clinicianid))
	sp.SetSafetyID(int(spuid))

	return sp
}

func (spd *SafetyPlanDao) GetByUserID(userID int) []*Model.SafetyPlan {
	var query = "SELECT safetyid, triggers, warningsigns, destructivebehaviors, internalstrategies, updateddatetime, updatedclinician, client_clientid, clinician_clinicianid FROM cfc.safety_plan join cfc.client on cfc.safety_plan.client_clientid = cfc.client.clientid WHERE cfc.client.person_userid=$1"
	var parameters = []interface{}{
		userID,
	}
	var spArray []*Model.SafetyPlan
	result, err := spd.db.Select(query, parameters)
	if err != nil || len(result) == 0 {
		return spArray
	}
	fmt.Println(result)
	for i := 0; i < len(result); i++ {
		spuid, _ := strconv.ParseInt(result[i][0], 10, 64)
		uc, _ := strconv.ParseInt(result[i][6], 10, 64)
		clientuid, _ := strconv.ParseInt(result[i][7], 10, 64)
		clinicianid, _ := strconv.ParseInt(result[i][8], 10, 64)
		sp := Model.NewSafetyPlan(result[i][1], result[i][2], result[i][3], result[i][4], result[i][5], int(uc), int(clientuid), int(clinicianid))
		sp.SetSafetyID(int(spuid))
		spArray = append(spArray, sp)
	}

	return spArray
}

func (spd *SafetyPlanDao) GetAll() ([]*Model.SafetyPlan, error) {
	var query = "SELECT * FROM cfc.safety_plan"
	var spList []*Model.SafetyPlan

	result, err := spd.db.Select(query, []interface{}{})
	if err != nil || len(result) == 0 {
		return spList, err
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

	return spList, nil
}

func (spd *SafetyPlanDao) Add(sp Model.SafetyPlan) (int, error) {
	var query = "INSERT INTO safety_plan(safetyId,triggers,warningSigns,destructiveBehaviors,internalStrategies,updatedDatetime,updatedClinician,Client_clientId,Clinician_clinicianId) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	var parameters = []interface{}{
		sp.GetSafetyID(),
		sp.GetTriggers(),
		sp.GetWarningSigns(),
		sp.GetDestructiveBehaviors(),
		sp.GetInternalStrategies(),
		sp.GetUpdatedDatetime(),
		sp.GetUpdatedClinician(),
		sp.GetClientID(),
		sp.GetClinicianID(),
	}

	return spd.db.Insert(query, parameters)
}

func (spd *SafetyPlanDao) Update(userID int, sp *Model.SafetyPlan) (int, error) {
	var query = "UPDATE safety_plan SET triggers=$1,warningSigns=$2,destructiveBehaviors=$3,internalStrategies=$4,updatedDatetime=$5,updatedClinician=$6,Client_clientId=$7,Clinician_clinicianId=$8 WHERE safetyId=$9"
	var parameters = []interface{}{
		sp.GetTriggers(),
		sp.GetWarningSigns(),
		sp.GetDestructiveBehaviors(),
		sp.GetInternalStrategies(),
		sp.GetUpdatedDatetime(),
		sp.GetUpdatedClinician(),
		sp.GetClientID(),
		sp.GetClinicianID(),
		userID,
	}

	return spd.db.Update(query, parameters)
}

func (spd *SafetyPlanDao) Delete(safetyId int) error {
	var query = "DELETE FROM safety_plan WHERE safetyId=$1"
	var parameters = []interface{}{
		safetyId,
	}

	return spd.db.Delete(query, parameters)
}

// TODO GetClientBySafetyPlanID()

// TODO GetClinicianBySafetyPlanID()
