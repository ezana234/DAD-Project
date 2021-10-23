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
	query := DB.NewNamedParameterQuery("SELECT * FROM clinician WHERE person.clinicianId=:clinicianID")
	var parameterMap = map[string]interface{}{
		"clinicianID": clinicianID,
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
	query := DB.NewNamedParameterQuery("SELECT * FROM clinician")
	var cList []*Model.Clinician

	result, err := cd.db.Select(query, map[string]interface{}{})
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
	query := DB.NewNamedParameterQuery("INSERT INTO clinician(Person_userId) VALUES(:userID)")
	var parameterMap = map[string]interface{}{
		"userID": c.UserID(),
	}

	err := cd.db.Update(query, parameterMap)

	return err
}

func (cd *ClinicianDao) UpdateClinician(clinicianID int, c *Model.Clinician) error {
	query := DB.NewNamedParameterQuery("UPDATE clinician SET Person_userId=:userID WHERE clinicianId=:clinicianID")
	var parameterMap = map[string]interface{}{
		"userID":      c.UserID(),
		"clinicianID": clinicianID,
	}

	err := cd.db.Update(query, parameterMap)

	return err
}

func (cd *ClinicianDao) DeleteClinician(clinicianID int) error {
	query := DB.NewNamedParameterQuery("DELETE FROM clinician WHERE clinicianId=:clinicianID")
	var parameterMap = map[string]interface{}{
		"clinicianID": clinicianID,
	}

	err := cd.db.Update(query, parameterMap)

	return err
}

func (cd *ClinicianDao) GetClientsByClinicianID(clinicianID int) []*Model.Client {
	query := DB.NewNamedParameterQuery("SELECT * FROM client WHERE clientId IN (SELECT Client_clientId FROM client_has_clinician WHERE Clinician_clinicianId=:clinicianID)")
	var parameterMap = map[string]interface{}{
		"clinicianID": clinicianID,
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

func (cd *ClinicianDao) GetNextID() int {
	query := DB.NewNamedParameterQuery("SELECT MAX(clinicianId) FROM clinician")

	result, err := cd.db.Select(query, map[string]interface{}{})
	if err != nil {
		return -1
	}

	res, _ := strconv.ParseInt(result[0][0], 10, 64)

	return int(res) + 1
}
