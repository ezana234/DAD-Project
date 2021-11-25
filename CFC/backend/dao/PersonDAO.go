package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
	"errors"
	"strconv"
)

type PersonDao struct {
	db DB.DatabaseConnection
}

func NewPersonDao(db DB.DatabaseConnection) *PersonDao {
	return &PersonDao{db: db}
}

func (pd *PersonDao) GetUserByID(userID int) (*Model.Person, error) {
	var query = "SELECT * FROM cfc.person WHERE userid=$1 LIMIT 1"
	var parameterMap = []interface{}{
		userID,
	}

	result, err := pd.db.Select(query, parameterMap)
	if err != nil {
		return new(Model.Person), err
	}

	var res = result[0]
	uid, _ := strconv.ParseInt(res[0], 10, 64)
	p := Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8], res[9], res[10])
	p.SetUserID(int(uid))

	return p, nil
}

func (pd *PersonDao) GetAll() ([]*Model.Person, error) {
	var query = "SELECT * FROM cfc.person"
	var pList []*Model.Person

	result, err := pd.db.Select(query, []interface{}{})
	if err != nil || len(result) == 0 {
		return pList, err
	}

	for _, res := range result {
		uid, _ := strconv.ParseInt(res[0], 10, 64)
		tmpP := Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8], res[9], res[10])
		tmpP.SetUserID(int(uid))
		pList = append(pList, tmpP)
	}

	return pList, nil
}

func (pd *PersonDao) Add(p Model.Person) (int, error) {
	var query = "INSERT INTO cfc.person(userid,username,password,firstname,lastname,email,address,phonenumber,role,expiration,dob) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
	var parameters = []interface{}{
		p.GetUserID(),
		p.GetUserName(),
		p.GetPassword(),
		p.GetFirstName(),
		p.GetLastName(),
		p.GetEmail(),
		p.GetAddress(),
		p.GetPhoneNumber(),
		p.GetRole(),
		p.GetExpiration(),
		p.GetDOB(),
	}

	return pd.db.Insert(query, parameters)
}

func (pd *PersonDao) Update(userID int, p *Model.Person) (int, error) {
	var query = "UPDATE cfc.person SET userName=$1, password=$2, firstName=$3, lastName=$4, email=$5, address=$6, phoneNumber=$7, role=$8, expiration=$9, dob=$10 WHERE userId=$11"
	var parameters = []interface{}{
		p.GetUserName(),
		p.GetPassword(),
		p.GetFirstName(),
		p.GetLastName(),
		p.GetEmail(),
		p.GetAddress(),
		p.GetPhoneNumber(),
		p.GetRole(),
		p.GetExpiration(),
		p.GetDOB(),
		userID,
	}

	return pd.db.Update(query, parameters)
}

func (pd *PersonDao) Delete(userID int) (int, error) {
	var query = "DELETE FROM cfc.person WHERE userId=$1"
	var parameters = []interface{}{
		userID,
	}

	return pd.db.Delete(query, parameters)
}

func (pd *PersonDao) GetPersonByUserName(userName string) (*Model.Person, error) {
	var query = "SELECT * FROM cfc.person WHERE userName=$1 LIMIT 1"
	var parameterMap = []interface{}{userName}

	result, err := pd.db.Select(query, parameterMap)
	if err != nil || len(result) == 0 {
		return new(Model.Person), err
	}

	var res = result[0]
	uid, _ := strconv.ParseInt(res[0], 10, 64)
	p := Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8], res[9], res[10])
	p.SetUserID(int(uid))

	return p, nil
}

// func (pd *PersonDao) GetPersonByEmail(email string, password string) (*Model.Person, error) {
// 	var query = "SELECT * FROM cfc.person WHERE email=$1 and password=$2 LIMIT 1"
// 	var parameterMap = []interface{}{email, password}

// 	result, err := pd.db.Select(query, parameterMap)
// 	if err != nil || len(result) == 0 {
// 		return new(Model.Person), err
// 	}

// 	var res = result[0]
// 	uid, _ := strconv.ParseInt(res[0], 10, 64)
// 	p := Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8], res[9], res[10])
// 	p.SetUserID(int(uid))

// 	return p, nil
// }

func (pd *PersonDao) GetPersonByEmail(email string) (*Model.Person, error) {
	var query = "SELECT * FROM cfc.person WHERE email=$1 LIMIT 1"
	var parameterMap = []interface{}{email}

	result, err := pd.db.Select(query, parameterMap)
	if err != nil || len(result) == 0 {
		return new(Model.Person), err
	}

	var res = result[0]
	uid, _ := strconv.ParseInt(res[0], 10, 64)
	p := Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8], res[9], res[10])
	p.SetUserID(int(uid))

	return p, nil
}

func (pd *PersonDao) GetNextUserID() int {
	var query = "SELECT MAX(userId) FROM cfc.person"

	result, err := pd.db.Select(query, []interface{}{})
	if err != nil {
		return -1
	}

	res, _ := strconv.ParseInt(result[0][0], 10, 64)

	return int(res) + 1
}

func (pd *PersonDao) UsernameExists(username string) (bool, error) {
	var query = "SELECT username FROM cfc.person WHERE username=$1"
	var parameters = []interface{}{username}

	result, err := pd.db.Select(query, parameters)
	if err != nil || len(result) > 0 {
		return true, err
	}

	return false, nil
}

func (pd *PersonDao) EmailExists(email string) (bool, error) {
	var query = "SELECT email FROM cfc.person WHERE email=$1"
	var parameters = []interface{}{email}

	result, err := pd.db.Select(query, parameters)
	if err != nil || len(result) > 0 {
		return true, err
	}

	return false, nil
}

// TODO GetClinicianByUserID()

func (pd *PersonDao) GetClientByUserID(userID int) (*Model.Client, error) {
	var query = "SELECT * FROM cfc.client WHERE userID=$1 LIMIT 1"
	var parameters = []interface{}{userID}

	result, err := pd.db.Select(query, parameters)
	if err != nil {
		return new(Model.Client), err
	}

	var res = result[0]
	cuid, _ := strconv.ParseInt(res[0], 10, 64)
	uid, _ := strconv.ParseInt(res[1], 10, 64)
	c := Model.NewClient(int(uid))
	c.SetClientID(int(cuid))

	return c, nil
}

func (pd *PersonDao) GetClinicianByUserID(userID int) (*Model.Clinician, error) {
	var query = "SELECT * FROM cfc.clinician WHERE clinician.person_userid=$1"
	var parameters = []interface{}{userID}

	result, err := pd.db.Select(query, parameters)
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

func (pd *PersonDao) GetSafetyPlansByUserID(userID int, role int) ([]*Model.SafetyPlan, error) {
	var spList []*Model.SafetyPlan
	var query string
	var parameters = []interface{}{userID}

	if role == 1 {
		query = "SELECT * FROM cfc.safety_plan WHERE safety_plan.client_clientid IN (SELECT clientid FROM cfc.client WHERE client.person_userid = $1)"
	} else if role == 2 {
		query = "SELECT * FROM cfc.safety_plan WHERE safety_plan.clinician_clinicianid IN (SELECT clinicianid FROM cfc.clinician WHERE person_userid = $1)"
	} else {
		return spList, errors.New("incorrect role id")
	}

	result, err := pd.db.Select(query, parameters)
	if err != nil {
		println(err)
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
	println(spList)
	return spList, nil
}

func (pd *PersonDao) GetAppointmentsByUserID(userID int, role int) ([]*Model.Appointment, error) {
	var aList []*Model.Appointment
	var query string
	var parameters = []interface{}{userID}

	if role == 1 {
		query = "SELECT * FROM cfc.appointments WHERE appointments.client_clientid IN (SELECT clientid FROM cfc.client WHERE client.person_userid=$1)"
	} else if role == 2 {
		query = "SELECT * FROM cfc.appointments WHERE appointments.clinician_clinicianid IN (SELECT clinicianid FROM cfc.clinician WHERE clinician.person_userid=$1)"
	} else {
		return aList, errors.New("incorrect role id")
	}

	result, err := pd.db.Select(query, parameters)
	if err != nil {
		return aList, err
	}

	for _, res := range result {
		appid, _ := strconv.ParseInt(res[0], 10, 64)
		clientid, _ := strconv.ParseInt(res[3], 10, 64)
		clinicianid, _ := strconv.ParseInt(res[4], 10, 64)
		app := Model.NewAppointment(res[1], res[2], int(clientid), int(clinicianid))
		app.SetAppointmentID(int(appid))

		aList = append(aList, app)
	}

	return aList, nil
}
