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
