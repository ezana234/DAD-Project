package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
	"strconv"
)

type ClinicianDao struct {
	db DB.DatabaseConnection
}

func NewClinicianDao(db DB.DatabaseConnection) *ClinicianDao {
	return &ClinicianDao{db: db}
}

func (cd *ClinicianDao) GetClinician(clinicianID int) *Model.Clinician {
	var query = "SELECT * FROM clinician WHERE person.clinicianId=$1"
	var parameterMap = []interface{}{
		clinicianID,
	}
	var c = new(Model.Clinician)

	result, err := cd.db.Select(query, parameterMap)
	if err != nil {
		return c
	}

	var res = result[0]

	cuid, _ := strconv.ParseInt(res[0], 10, 64)
	uid, _ := strconv.ParseInt(res[1], 10, 64)
	c = Model.NewClinician(int(uid))
	c.SetClinicianID(int(cuid))

	return c
}

func (cd *ClinicianDao) GetAllClinicians() []*Model.Clinician {
	var query = "SELECT * FROM clinician"
	var cList []*Model.Clinician

	result, err := cd.db.Select(query, []interface{}{})
	if err != nil {
		return cList
	}

	for _, res := range result {
		cuid, _ := strconv.ParseInt(res[0], 10, 64)
		uid, _ := strconv.ParseInt(res[1], 10, 64)
		tmpC := Model.NewClinician(int(uid))
		tmpC.SetClinicianID(int(cuid))
		cList = append(cList, tmpC)
	}

	return cList
}

func (cd *ClinicianDao) AddClinician(c Model.Clinician) error {
	var query = "INSERT INTO clinician(Person_userId) VALUES($1)"
	var parameters = []interface{}{
		c.GetUserID(),
	}

	err := cd.db.Insert(query, parameters)

	return err
}

func (cd *ClinicianDao) UpdateClinician(clinicianID int, c *Model.Clinician) error {
	var query = "UPDATE clinician SET Person_userId=$1 WHERE clinicianId=$2"
	var parameters = []interface{}{
		c.GetUserID(),
		clinicianID,
	}

	err := cd.db.Update(query, parameters)

	return err
}

func (cd *ClinicianDao) DeleteClinician(clinicianID int) error {
	var query = "DELETE FROM clinician WHERE clinicianId=$1"
	var parameters = []interface{}{
		clinicianID,
	}

	err := cd.db.Delete(query, parameters)

	return err
}

func (cd *ClinicianDao) GetClientsByClinicianID(clinicianID int) []*Model.Client {
	var query = "SELECT * FROM client WHERE clientId IN (SELECT Client_clientId FROM client_has_clinician WHERE Clinician_clinicianId=$1)"
	var parameterMap = []interface{}{
		clinicianID,
	}

	var cList []*Model.Client

	result, err := cd.db.Select(query, parameterMap)
	if err != nil {
		return cList
	}

	for _, res := range result {
		cuid, _ := strconv.ParseInt(res[0], 10, 64)
		uid, _ := strconv.ParseInt(res[1], 10, 64)
		tmpC := Model.NewClient(int(uid))
		tmpC.SetClientID(int(cuid))
		cList = append(cList, tmpC)
	}

	return cList
}

func (cd *ClinicianDao) GetAppointmentsByClinicianID(clinicianID int) []*Model.Appointment {
	var query = "SELECT * FROM appointments WHERE Clinician_clinicianId=$1"
	var parameterMap = []interface{}{
		clinicianID,
	}

	var aList []*Model.Appointment

	result, err := cd.db.Select(query, parameterMap)
	if err != nil {
		return aList
	}

	for _, res := range result {
		aid, _ := strconv.ParseInt(res[0], 10, 64)
		clieID, _ := strconv.ParseInt(res[3], 10, 64)
		clinID, _ := strconv.ParseInt(res[4], 10, 64)

		a := Model.NewAppointment(res[1], res[2], int(clieID), int(clinID))
		a.SetAppointmentID(int(aid))
		aList = append(aList, a)
	}

	return aList
}

func (cd *ClinicianDao) GetSafetyPlansByClinicianID(clinicianID int) []*Model.SafetyPlan {
	var query = "SELECT * FROM safety_plan WHERE Clinician_clinicianId=$1"
	var parameterMap = []interface{}{
		clinicianID,
	}

	var spList []*Model.SafetyPlan

	result, err := cd.db.Select(query, parameterMap)
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

func (cd *ClinicianDao) GetNextClinicianID() int {
	var query = "SELECT MAX(clinicianId) FROM clinician"

	result, err := cd.db.Select(query, []interface{}{})
	if err != nil {
		return -1
	}

	res, _ := strconv.ParseInt(result[0][0], 10, 64)

	return int(res) + 1
}
