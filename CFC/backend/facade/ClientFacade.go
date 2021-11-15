package facade

import (
	"CFC/backend/CFC/backend/DB"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
	"log"
)

type ClientFacade struct {
	clientDao DAO.ClientDao
}

func NewClientFacade(db DB.DatabaseConnection) *ClientFacade {
	return &ClientFacade{clientDao: *DAO.NewClientDao(db)}
}

//func (cf *ClientFacade) GetClient(clientID int) *Model.Client {
//	return cf.clientDao.GetClient(clientID)
//}
//
//func (cf *ClientFacade) AddClient(c Model.Client) interface{} {
//	_ = cf.clientDao.AddClient(c)
//	return nil
//}

func (cf *ClientFacade) GetClientByID(clientID int) (*Model.Client, int) {
	c, err := cf.clientDao.GetClientByID(clientID)

	if err != nil {
		return new(Model.Client), 0
	}

	return c, 0
}

func (cf *ClientFacade) AddClient(c Model.Client) int {
	c.SetClientID(cf.clientDao.GetNextClientID())

	err := cf.clientDao.Add(c)
	if err != nil {
		log.Printf("Error: %s when adding person", err)
		return 0
	}

	return 1
}

func (cf *ClientFacade) GetSafetyPlanByClientID(clientID int) (*Model.SafetyPlan, int) {
	sp, err := cf.clientDao.GetSafetyPlanByClientID(clientID)
	if err != nil {
		return new(Model.SafetyPlan), 0
	}

	return sp, 1
}
