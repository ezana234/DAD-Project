package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
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

func (pd *PersonDao) AddPerson(p Model.Person) error {
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

	if err != nil {
		return err
	}
	return nil
}

func (pd *PersonDao) DeletePersonByID(userID int) error {
	query := DB.NewNamedParameterQuery("DELETE FROM person WHERE userId=:userID")
	var parameterMap = map[string]interface{}{
		"userID": userID,
	}

	err := pd.db.Update(query, parameterMap)
	if err != nil {
		return err
	}
	return nil
}
