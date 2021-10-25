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
	return pf.personDao.GetByID(userID)
}

func (pf *PersonFacade) GetPersonByEmail(email string, password string) *Model.Person {
	return pf.personDao.GetByEmail(email, password)
}

func (pf *PersonFacade) AddPerson(p Model.Person) error {
	//p.SetUserID(pf.personDao.)
	//_ = pf.personDao.Add(p)
	return nil
}

func (pf *PersonFacade) DeletePerson(userID int) error {
	_ = pf.personDao.Delete(userID)
	return nil
}
