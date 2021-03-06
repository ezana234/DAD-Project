package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
	"strconv"
)

type ClinicianDao struct {
	db DB.DBConnection
}

func NewClinicianDao(db DB.DBConnection) *ClinicianDao {
	return &ClinicianDao{db: db}
}

func (cd *ClinicianDao) GetClinicianByID(clinicianID int) (*Model.Clinician, error) {
	var query = "SELECT * FROM cfc.clinician WHERE person.clinicianId=$1"
	var parameters = []interface{}{clinicianID}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return new(Model.Clinician), err
	}

	var res = result[0]
	cuid, _ := strconv.ParseInt(res[0], 10, 64)
	uid, _ := strconv.ParseInt(res[1], 10, 64)
	c := Model.NewClinician(int(uid), res[2])
	c.SetClinicianID(int(cuid))

	return c, nil
}

func (cd *ClinicianDao) GetAllClinicians() ([]*Model.Clinician, error) {
	var query = "SELECT * FROM cfc.clinician"
	var cList []*Model.Clinician

	result, err := cd.db.Select(query, []interface{}{})
	if err != nil {
		return cList, err
	}

	for _, res := range result {
		cuid, _ := strconv.ParseInt(res[0], 10, 64)
		uid, _ := strconv.ParseInt(res[1], 10, 64)
		tmpC := Model.NewClinician(int(uid), res[2])
		tmpC.SetClinicianID(int(cuid))
		cList = append(cList, tmpC)
	}

	return cList, nil
}

func (cd *ClinicianDao) GetAllClients() ([]*Model.Person, error) {
	var query = "select * from cfc.client join cfc.person on person.userid = client.person_userid;"
	var pList []*Model.Person

	result, err := cd.db.Select(query, []interface{}{})
	if err != nil || len(result) == 0 {
		return pList, err
	}

	for _, res := range result {
		//fmt.Println(res[8])
		uid, _ := strconv.ParseInt(res[2], 10, 64)
		tmpP := Model.NewPerson(res[3], res[4], res[5], res[6], res[7], res[8], res[9], res[10], res[11], res[12])
		tmpP.SetUserID(int(uid))
		pList = append(pList, tmpP)
	}

	return pList, nil
}

func (cd *ClinicianDao) AddClinician(c Model.Clinician) (int, error) {
	var query = "INSERT INTO cfc.clinician(clinicianid,Person_userId,referral) VALUES($1,$2,$3)"
	var parameters = []interface{}{
		cd.GetNextClinicianID(),
		c.GetUserID(),
		c.GetReferral(),
	}

	return cd.db.Insert(query, parameters)
}

func (cd *ClinicianDao) UpdateClinician(clinicianID int, c *Model.Clinician) (int, error) {
	var query = "UPDATE cfc.clinician SET Person_userId=$1, referral=$2 WHERE clinicianId=$3"
	var parameters = []interface{}{
		c.GetUserID(),
		c.GetReferral(),
		clinicianID,
	}

	return cd.db.Update(query, parameters)
}

func (cd *ClinicianDao) DeleteClinician(clinicianID int) (int, error) {
	var query = "DELETE FROM cfc.clinician WHERE clinicianId=$1"
	var parameters = []interface{}{
		clinicianID,
	}

	return cd.db.Delete(query, parameters)
}

func (cd *ClinicianDao) GetUserByClinicianID(clinicianID int) (*Model.Person, error) {
	var query = "SELECT * FROM cfc.person WHERE person.userid IN (SELECT person_userid FROM cfc.clinician WHERE clinician.clinicianid = $1)"
	var parameters = []interface{}{clinicianID}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return new(Model.Person), err
	}

	var res = result[0]
	uid, _ := strconv.ParseInt(res[0], 10, 64)
	p := Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8], res[9], res[10])
	p.SetUserID(int(uid))

	return p, nil
}

func (cd *ClinicianDao) GetClientsByClinicianID(clinicianID int) []*Model.Client {
	var query = "SELECT * FROM cfc.client WHERE clientId IN (SELECT Client_clientId FROM cfc.client_has_clinician WHERE Clinician_clinicianId=$1) ORDER BY client.clientID"
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

func (cd *ClinicianDao) GetClientUsersByClinicianID(clinicianID int) ([]*Model.Person, error) {
	var pList []*Model.Person
	var query = "SELECT * FROM cfc.person WHERE person.userid IN (SELECT person_userID FROM cfc.client INNER JOIN cfc.client_has_clinician ON client.clientid=client_has_clinician.client_clientid AND client_has_clinician.clinician_clinicianid=$1) ORDER BY person.userid"
	var parameters = []interface{}{clinicianID}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return pList, err
	}

	for _, res := range result {
		uid, _ := strconv.ParseInt(res[0], 10, 64)
		p := Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8], res[9], res[10])
		p.SetUserID(int(uid))
		pList = append(pList, p)
	}

	return pList, nil
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

func (cd *ClinicianDao) GetSafetyPlansByClinicianID(clinicianID int) ([]*Model.SafetyPlan, error) {
	var query = "SELECT * FROM safety_plan WHERE Clinician_clinicianId=$1"
	var parameterMap = []interface{}{
		clinicianID,
	}

	var spList []*Model.SafetyPlan

	result, err := cd.db.Select(query, parameterMap)
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

func (cd *ClinicianDao) GetNextClinicianID() int {
	var query = "SELECT MAX(clinicianId) FROM clinician"

	result, err := cd.db.Select(query, []interface{}{})
	if err != nil {
		return -1
	}

	res, _ := strconv.ParseInt(result[0][0], 10, 64)

	return int(res) + 1
}

func (cd *ClinicianDao) GetClinicianByUserID(userID int) (*Model.Clinician, error) {
	var query = "SELECT * FROM cfc.clinician WHERE Person_userId=:$1"
	var parameters = []interface{}{userID}

	result, err := cd.db.Select(query, parameters)
	if err != nil || len(result) == 0 {
		return new(Model.Clinician), err
	}

	var res = result[0]
	uid, _ := strconv.ParseInt(res[0], 10, 64)
	cuid, _ := strconv.ParseInt(res[1], 10, 64)
	clinician := Model.NewClinician(int(uid), res[2])
	clinician.SetClinicianID(int(cuid))

	return clinician, nil
}

func (cd *ClinicianDao) GetClinicianByReferral(referral string) (*Model.Clinician, error) {
	var query = "SELECT * FROM cfc.clinician WHERE referral=$1"
	var parameters = []interface{}{referral}

	result, err := cd.db.Select(query, parameters)
	if err != nil || len(result) == 0 {
		return new(Model.Clinician), err
	}

	var res = result[0]
	cuid, _ := strconv.ParseInt(res[0], 10, 64)
	uid, _ := strconv.ParseInt(res[1], 10, 64)
	clinician := Model.NewClinician(int(uid), res[2])
	clinician.SetClinicianID(int(cuid))

	return clinician, nil
}

func (cd *ClinicianDao) GetClinicianNameByClinicianID(clinicianID int) (*Model.Person, error) {
	var query = "SELECT userid, firstname, lastname FROM cfc.\"ClinicianNames\" WHERE clinicianid = $1"
	var parameters = []interface{}{clinicianID}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return new(Model.Person), err
	}

	uid, _ := strconv.ParseInt(result[0][0], 10, 64)
	clinicianName := Model.NewPerson("", "", result[0][1], result[0][2], "", "", "", "", "", "")
	clinicianName.SetUserID(int(uid))

	return clinicianName, nil
}

func (cd *ClinicianDao) GetAllClinicianNames() ([]*Model.Person, error) {
	var clinicianNames []*Model.Person
	var query = "SELECT userid, firstname, lastname FROM cfc.\"ClinicianNames\""
	var parameters []interface{}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return clinicianNames, err
	}

	for _, res := range result {
		uid, _ := strconv.ParseInt(res[0], 10, 64)
		clinicianName := Model.NewPerson("", "", res[1], res[2], "", "", "", "", "", "")
		clinicianName.SetUserID(int(uid))
		clinicianNames = append(clinicianNames, clinicianName)
	}

	return clinicianNames, nil
}

// TODO GetUserByClinicianID()

// TODO GetClientsByClinicianID()

// TODO GetAppointmentsByClinicianID()

// TODO GetSafetyPlansByClinicianID()
