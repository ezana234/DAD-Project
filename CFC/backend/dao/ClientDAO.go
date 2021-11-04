package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
	"strconv"
)

type ClientDao struct {
	db DB.DatabaseConnection
}

func NewClientDao(db DB.DatabaseConnection) *ClientDao {
	return &ClientDao{db: db}
}

func (cd *ClientDao) GetClient(clientID int) *Model.Client {
	query := DB.NewNamedParameterQuery("SELECT * FROM client WHERE person.clientId=:clientID")
	var parameterMap = map[string]interface{}{
		"clientID": clientID,
	}
	var c = new(Model.Client)

	result, err := cd.db.Select(query, parameterMap)
	if err != nil {
		return c
	}

	var res = result[0]
	cuid, _ := strconv.ParseInt(res[0], 10, 64)
	uid, _ := strconv.ParseInt(res[1], 10, 64)
	c = Model.NewClient(int(uid))
	c.SetClientID(int(cuid))

	return c
}

func (cd *ClientDao) AddClient(c Model.Client) error {
	return nil
}

func (cd *ClientDao) UpdateClient(c Model.Client) error {
	return nil
}

func (cd *ClientDao) DeleteClient(clientID int) error {
	return nil
}

func (cd *ClientDao) GetAllClients() error {
	return nil
}

// TODO GetSafetyPlanByClientID()

// TODO GetFamilyMemberByClientID()

// TODO GetSupportNetworkByClientID()

// TODO GetClinicianByClientID()

// TODO GetAppointmentByClientID()

// TODO GetPersonByClientID()
