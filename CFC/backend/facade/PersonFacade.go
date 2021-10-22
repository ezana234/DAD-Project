package facade

import (
	"CFC/backend/CFC/backend/DB"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
)

type PersonFacade struct {
	personDao DAO.PersonDao
}

func NewPersonFacade(db DB.DatabaseConnection) *PersonFacade {
	return &PersonFacade{personDao: *DAO.NewPersonDao(db)}
}

func (pf *PersonFacade) GetPerson(userID int) *Model.Person {
	var p = pf.personDao.GetByID(userID)
	return p
}

func (pf *PersonFacade) AddPerson(p Model.Person) interface{} {
	_ = pf.personDao.AddPerson(p)
	return nil
}

func (pf *PersonFacade) DeletePerson(userID int) interface{} {
	_ = pf.personDao.DeletePersonByID(userID)
	return nil
}
