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

func (cf *ClientFacade) GetClientByClientID(clientID int) (*Model.Client, int) {
	c, err := cf.clientDao.GetClientByID(clientID)

	if err != nil {
		return new(Model.Client), 0
	}

	return c, 0
}

func (cf *ClientFacade) GetAllClients() ([]*Model.Client, int) {
	var emptyList []*Model.Client
	cList, err := cf.clientDao.GetAll()
	if err != nil {
		return emptyList, 0
	}

	return cList, 1
}

func (cf *ClientFacade) AddClient(c Model.Client) int {
	rowsAffected, err := cf.clientDao.Add(c)
	if err != nil {
		log.Printf("Error: %s when adding client", err)
		return 0
	}

	if rowsAffected == 0 {
		log.Printf("0 rows affected when adding client")
		return 0
	}

	return 1
}

func (cf *ClientFacade) DeleteClient(clientID int) int {
	err := cf.clientDao.Delete(clientID)
	if err != nil {
		return 0
	}

	return 1
}

func (cf *ClientFacade) UpdateClient(clientID int, c *Model.Client) int {
	rowsAffected, err := cf.clientDao.Update(clientID, c)
	if err != nil {
		return 0
	}

	if rowsAffected == 0 {
		return -1
	}

	return 1
}

func (cf *ClientFacade) GetUserByClientID(clientID int) (*Model.Person, int) {
	p, err := cf.clientDao.GetUserByClientID(clientID)
	if err != nil {
		return new(Model.Person), 0
	}

	return p, 1
}

// func (cf *ClientFacade) GetSafetyPlanByClientID(clientID int) (*Model.SafetyPlan, int) {
// 	sp, err := cf.clientDao.GetSafetyPlanByClientID(clientID)
// 	if err != nil {
// 		return new(Model.SafetyPlan), 0
// 	}

// 	return sp, 1
// }

func (cf *ClientFacade) GetClinicianByClientID(clientID int) (*Model.Clinician, int) {
	clinician, err := cf.clientDao.GetClinicianByClientID(clientID)
	if err != nil {
		return new(Model.Clinician), 0
	}

	return clinician, 1
}

func (cf *ClientFacade) GetUserClinicianByClientID(clientID int) (*Model.Person, int) {
	p, err := cf.clientDao.GetClinicianUserByClientID(clientID)
	if err != nil {
		return new(Model.Person), 0
	}

	return p, 1
}
