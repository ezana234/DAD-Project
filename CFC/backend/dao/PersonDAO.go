package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
	"fmt"
	"strconv"
)

type PersonDao struct {
	db DB.DatabaseConnection
}

func NewPersonDao(db DB.DatabaseConnection) *PersonDao {
	return &PersonDao{db: db}
}

func (pd *PersonDao) GetByID(userID int) *Model.Person {
	query := DB.NewNamedParameterQuery("SELECT * FROM person WHERE person.UserId=:userID")
	var parameterMap = map[string]interface{}{
		"userID": userID,
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

func (pd *PersonDao) GetByEmail(email string, password string) *Model.Person {
	fmt.Println(email, password)
	query := DB.NewNamedParameterQuery("SELECT * FROM person WHERE person.Email=:email AND person.Password=:password;")
	var parameterMap = map[string]interface{}{
		"email":    email,
		"password": password}

	var p = new(Model.Person)

	result, err := pd.db.Select(query, parameterMap)
	if err != nil || len(result) == 0 {
		return p
	}

	var res = result[0]

	uid, _ := strconv.ParseInt(res[0], 10, 64)
	p = Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8])
	p.SetUserID(int(uid))

	return p
}

func (pd *PersonDao) GetAll() []*Model.Person {
	query := DB.NewNamedParameterQuery("SELECT * FROM person")
	var pList []*Model.Person

	result, err := pd.db.Select(query, map[string]interface{}{})
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
	query := DB.NewNamedParameterQuery("INSERT INTO person(userId,userName,password,firstName,lastName,email,address,phoneNumber,role) VALUES(:userID,:userName,:password,:firstName,:lastName,:email,:address,:phoneNumber,:role)")
	var parameterMap = map[string]interface{}{
		"userID":      p.GetUserID(),
		"userName":    p.GetUserName(),
		"password":    p.GetPassword(),
		"firstName":   p.GetFirstName(),
		"lastName":    p.GetLastName(),
		"email":       p.GetEmail(),
		"address":     p.GetAddress(),
		"phoneNumber": p.GetPhoneNumber(),
		"role":        p.GetRole(),
	}

	return pd.db.Update(query, parameterMap)
}

func (pd *PersonDao) Update(userID int, p *Model.Person) error {
	query := DB.NewNamedParameterQuery("UPDATE person SET userName=:userName, password=:password, firstName=:firstName, lastName=:lastName, email=:email, address=:address, phoneNumber=:phoneNumber, role=:role WHERE userId=:userID")
	var parameterMap = map[string]interface{}{
		"userName":    p.GetUserName(),
		"password":    p.GetPassword(),
		"firstName":   p.GetFirstName(),
		"lastName":    p.GetLastName(),
		"email":       p.GetEmail(),
		"address":     p.GetAddress(),
		"phoneNumber": p.GetPhoneNumber(),
		"role":        p.GetRole(),
		"userID":      userID,
	}

	return pd.db.Update(query, parameterMap)
}

func (pd *PersonDao) Delete(userID int) error {
	query := DB.NewNamedParameterQuery("DELETE FROM person WHERE userId=:userID")
	var parameterMap = map[string]interface{}{
		"userID": userID,
	}

	return pd.db.Update(query, parameterMap)
}

func (pd *PersonDao) GetPersonsByUserName(userName string) []*Model.Person {
	query := DB.NewNamedParameterQuery("SELECT * FROM person WHERE userName=:userName")
	var parameterMap = map[string]interface{}{
		"userName": userName,
	}

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
	query := DB.NewNamedParameterQuery("SELECT userId, userName, password, email, role FROM person WHERE email=:email")
	var parameterMap = map[string]interface{}{
		"email": email,
	}

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
	query := DB.NewNamedParameterQuery("SELECT * FROM clinician WHERE Person_userId=:userID")
	var parameterMap = map[string]interface{}{
		"userID": userID,
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
	query := DB.NewNamedParameterQuery("SELECT * FROM client WHERE Person_userId=:userID")
	var parameterMap = map[string]interface{}{
		"userID": userID,
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
	query := DB.NewNamedParameterQuery("SELECT MAX(userId) FROM person")

	result, err := pd.db.Select(query, map[string]interface{}{})
	if err != nil {
		return -1
	}

	res, _ := strconv.ParseInt(result[0][0], 10, 64)

	return int(res) + 1
}
