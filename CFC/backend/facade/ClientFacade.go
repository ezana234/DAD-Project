package facade

import (
	"CFC/backend/CFC/backend/DB"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
)

type ClientFacade struct {
	clientDao DAO.ClientDao
}

func NewClientFacade(db DB.DatabaseConnection) *ClientFacade {
	return &ClientFacade{clientDao: *DAO.NewClientDao(db)}
}

func (cf *ClientFacade) GetClient(clientID int) *Model.Client {
	return cf.clientDao.GetClient(clientID)
}

func (cf *ClientFacade) AddClient(c Model.Client) interface{} {
	_ = cf.clientDao.AddClient(c)
	return nil
}
