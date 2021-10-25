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

	var p = new(Model.Person)

	result, err := pd.db.Select(query, parameterMap)
	if err != nil {
		return p
	}

	var res = result[0]

	uid, _ := strconv.ParseInt(res[0], 10, 64)
	p = Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8])
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
	if err != nil {
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
		"userID":      p.UserID(),
		"userName":    p.UserName(),
		"password":    p.Password(),
		"firstName":   p.FirstName(),
		"lastName":    p.LastName(),
		"email":       p.Email(),
		"address":     p.Address(),
		"phoneNumber": p.PhoneNumber(),
		"role":        p.Role(),
	}

	err := pd.db.Update(query, parameterMap)

	return err
}

func (pd *PersonDao) Update(userID int, p *Model.Person) error {

	//query := DB.NewNamedParameterQuery("UPDATE person SET userName=:userName, password=:password, firstName=:firstName, lastName=:lastName, email=:email, address=:address, phoneNumber=:phoneNumber, role=:role WHERE userId=:userID")
	//var parameterMap = map[string]interface{}{
	//	"userName":    p.UserName(),
	//	"password":    p.Password(),
	//	"firstName":   p.FirstName(),
	//	"lastName":    p.LastName(),
	//	"email":       p.Email(),
	//	"address":     p.Address(),
	//	"phoneNumber": p.PhoneNumber(),
	//	"role":        p.Role(),
	//}
	return nil
}

func (pd *PersonDao) Delete(userID int) error {
	query := DB.NewNamedParameterQuery("DELETE FROM person WHERE userId=:userID")
	var parameterMap = map[string]interface{}{
		"userID": userID,
	}

	err := pd.db.Update(query, parameterMap)

	return err
}
