package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strconv"
)

type PersonDao struct {
	db DB.DatabaseConnection
}

func NewPersonDao(db DB.DatabaseConnection) *PersonDao {
	return &PersonDao{db: db}
}

func (pd *PersonDao) GetByID(userID int) *Model.Person {
	var query = "SELECT * FROM cfc.person WHERE userid=$1"
	var parameterMap = []interface{}{
		userID,
	}

	result, err := pd.db.Select(query, parameterMap)
	if err != nil {
		return new(Model.Person)
	}

	uid, _ := strconv.ParseInt(result[0][0], 10, 64)
	p := Model.NewPerson(result[0][1], result[0][2], result[0][3], result[0][4], result[0][5], result[0][6], result[0][7], result[0][8])
	p.SetUserID(int(uid))

	return p
}

func (pd *PersonDao) GetAll() []*Model.Person {
	var query = "SELECT * FROM cfc.person"
	var pList []*Model.Person

	result, err := pd.db.Select(query, []interface{}{})
	if err != nil || len(result) == 0 {
		return pList
	}

	for _, res := range result {
		uid, _ := strconv.ParseInt(res[0], 10, 64)
		tmpP := Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8])
		tmpP.SetUserID(int(uid))
		pList = append(pList, tmpP)
	}

	return pList
}

func (pd *PersonDao) Add(p Model.Person) error {
	var query = "INSERT INTO cfc.person(userid,username,password,firstname,lastname,email,address,phonenumber,role) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
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
	}

	return pd.db.Insert(query, parameters)
}

func (pd *PersonDao) Update(userID int, p *Model.Person) error {
	var query = "UPDATE cfc.person SET userName=$1, password=$2, firstName=$3, lastName=$4, email=$5, address=$6, phoneNumber=$7, role=$8 WHERE userId=$9"
	var parameters = []interface{}{
		p.GetUserName(),
		p.GetPassword(),
		p.GetFirstName(),
		p.GetLastName(),
		p.GetEmail(),
		p.GetAddress(),
		p.GetPhoneNumber(),
		p.GetRole(),
		userID,
	}

	return pd.db.Update(query, parameters)
}

func (pd *PersonDao) Delete(userID int) error {
	var query = "DELETE FROM cfc.person WHERE userId=$1"
	var parameters = []interface{}{
		userID,
	}

	return pd.db.Delete(query, parameters)
}

func (pd *PersonDao) GetPersonsByUserName(userName string) []*Model.Person {
	var query = "SELECT * FROM cfc.person WHERE userName=$1"
	var parameterMap = []interface{}{userName}

	var pList []*Model.Person

	result, err := pd.db.Select(query, parameterMap)
	if err != nil || len(result) == 0 {
		return pList
	}

	for _, res := range result {
		uid, _ := strconv.ParseInt(res[0], 10, 64)
		tmpP := Model.NewPerson(res[1], res[2], "", "", res[5], "", "", res[8])
		tmpP.SetUserID(int(uid))
		pList = append(pList, tmpP)
	}

	return pList
}

func (pd *PersonDao) GetPersonsByEmail(email string) []*Model.Person {
	var query = "SELECT userid, username, password, email, role FROM cfc.person WHERE email=$1"
	var parameterMap = []interface{}{email}

	var pList []*Model.Person

	result, err := pd.db.Select(query, parameterMap)
	if err != nil || len(result) == 0 {
		return pList
	}

	for _, res := range result {
		uid, _ := strconv.ParseInt(res[0], 10, 64)
		tmpP := Model.NewPerson(res[1], res[2], "", "", res[5], "", "", res[8])
		tmpP.SetUserID(int(uid))
		pList = append(pList, tmpP)
	}

	return pList
}

func (pd *PersonDao) GetClinicianByUserID(userID int) *Model.Clinician {
	var query = "SELECT * FROM cfc.clinician WHERE Person_userId=:$1"
	var parameterMap = []interface{}{
		userID,
	}

	result, err := pd.db.Select(query, parameterMap)
	if err != nil || len(result) == 0 {
		return new(Model.Clinician)
	}

	uid, _ := strconv.ParseInt(result[0][0], 10, 64)
	cuid, _ := strconv.ParseInt(result[0][1], 10, 64)
	clinician := Model.NewClinician(int(cuid))
	clinician.SetClinicianID(int(uid))

	return clinician
}

func (pd *PersonDao) GetClientByUserID(userID int) *Model.Client {
	var query = "SELECT * FROM cfc.client WHERE Person_userId=$1"
	var parameterMap = []interface{}{
		userID,
	}

	result, err := pd.db.Select(query, parameterMap)
	if err != nil || len(result) == 0 {
		return new(Model.Client)
	}

	uid, _ := strconv.ParseInt(result[0][0], 10, 64)
	cuid, _ := strconv.ParseInt(result[0][1], 10, 64)
	c := Model.NewClient(int(cuid))
	c.SetClientID(int(uid))

	return c
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
