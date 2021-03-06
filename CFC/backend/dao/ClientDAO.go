package dao

import (
	"CFC/backend/CFC/backend/DB"
	Model "CFC/backend/CFC/backend/model"
	"strconv"
	"strings"
)

type ClientDao struct {
	db DB.DBConnection
}

func NewClientDao(db DB.DBConnection) *ClientDao {
	return &ClientDao{db: db}
}

func (cd *ClientDao) GetClientByID(clientID int) (*Model.Client, error) {
	var query = "SELECT * FROM cfc.client WHERE client.clientId=$1"
	var parameters = []interface{}{clientID}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return new(Model.Client), err
	}

	var res = result[0]
	cuid, _ := strconv.ParseInt(res[0], 10, 64)
	uid, _ := strconv.ParseInt(res[1], 10, 64)
	c := Model.NewClient(int(uid))
	c.SetClientID(int(cuid))

	return c, nil
}

func (cd *ClientDao) GetAll() ([]*Model.Client, error) {
	var query = "SELECT * FROM cfc.client"
	var cList []*Model.Client

	result, err := cd.db.Select(query, []interface{}{})
	if err != nil || len(result) == 0 {
		return cList, err
	}

	for _, res := range result {
		cuid, _ := strconv.ParseInt(res[0], 10, 64)
		uid, _ := strconv.ParseInt(res[1], 10, 64)
		tmpC := Model.NewClient(int(uid))
		tmpC.SetClientID(int(cuid))
		cList = append(cList, tmpC)
	}

	return cList, nil
}

func (cd *ClientDao) Add(c Model.Client) (int, error) {
	var query = "INSERT INTO cfc.client(clientid,person_userid) VALUES($1,$2);"
	var parameters = []interface{}{cd.GetNextClientID(), c.GetUserID()}

	return cd.db.Insert(query, parameters)
}

func (cd *ClientDao) Update(clientID int, c *Model.Client) (int, error) {
	var query = "UPDATE cfc.client SET userid=$1 WHERE clientid=$2"
	var parameters = []interface{}{c.GetUserID(), clientID}

	return cd.db.Update(query, parameters)
}

func (cd *ClientDao) Delete(clientID int) (int, error) {
	var query = "DELETE FROM cfc.client WHERE clientid=$1"
	var parameters = []interface{}{clientID}

	return cd.db.Delete(query, parameters)
}

func (cd *ClientDao) GetNextClientID() int {
	var query = "SELECT MAX(clientId) FROM cfc.client"

	result, err := cd.db.Select(query, []interface{}{})
	if err != nil {
		return -1
	}

	res, _ := strconv.ParseInt(result[0][0], 10, 64)

	return int(res) + 1
}

// TODO GetSupportNetworksByClientID()

// TODO GetClinicianByClientID()

// TODO GetAppointmentsByClientID()

func (cd *ClientDao) GetUserByClientID(clientID int) (*Model.Person, error) {
	var query = "SELECT * FROM cfc.person WHERE person.userid IN (SELECT person_userid FROM cfc.client WHERE clientid=$1);"
	var parameters = []interface{}{clientID}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return new(Model.Person), err
	}

	var res = result[0]
	uid, _ := strconv.ParseInt(res[0], 10, 64)
	p := Model.NewPerson(res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8], res[9], res[10])
	p.SetUserID(int(uid))

	return p, nil
}

// func (cd *ClientDao) GetSafetyPlanByClientID(clientID int) (*Model.SafetyPlan, error) {
// 	var query = "SELECT * FROM cfc.safety_plan WHERE clientid=$1"
// 	var parameters = []interface{}{clientID}

// 	result, err := cd.db.Select(query, parameters)
// 	if err != nil {
// 		return new(Model.SafetyPlan), err
// 	}

// 	var res = result[0]
// 	spuid, _ := strconv.ParseInt(res[0], 10, 64)
// 	uc, _ := strconv.ParseInt(res[6], 10, 64)
// 	clientuid, _ := strconv.ParseInt(res[7], 10, 64)
// 	clinicianid, _ := strconv.ParseInt(res[8], 10, 64)
// 	sp := Model.NewSafetyPlan(res[1], res[2], res[3], res[4], res[5], int(uc), int(clientuid), int(clinicianid))
// 	sp.SetSafetyID(int(spuid))

// 	return sp, nil
// }

func (cd *ClientDao) GetClinicianUserByClientID(clientID int) (*Model.Person, error) {
	var query = "SELECT * FROM cfc.person WHERE person.userid IN (SELECT clinician.person_userid FROM cfc.clinician INNER JOIN client_has_clinician ON client_has_clinician.clinician_clinicianid = clinician.clinicianid AND client_has_clinician.client_clientid=$1)"
	var parameters = []interface{}{clientID}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return new(Model.Person), err
	}

	var res = result[0]
	uid, _ := strconv.ParseInt(res[0], 10, 64)
	p := Model.NewPerson(res[1], res[2], res[3], res[4], strings.ToLower(res[5]), res[6], res[7], res[8], res[9], res[10])
	p.SetUserID(int(uid))

	return p, nil
}

func (cd *ClientDao) GetClinicianByClientID(clientID int) (*Model.Clinician, error) {
	var query = "SELECT * FROM cfc.clinician WHERE clinician.clinicianid IN (SELECT clinician_clinicianID FROM cfc.client_has_clinician WHERE client_has_clinician.client_clientid=$1)"
	var parameters = []interface{}{clientID}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return new(Model.Clinician), err
	}

	var res = result[0]
	cuid, _ := strconv.ParseInt(res[0], 10, 64)
	uid, _ := strconv.ParseInt(res[1], 10, 64)
	c := Model.NewClinician(int(uid), res[2])
	c.SetClinicianID(int(cuid))

	return c, nil
}

func (cd *ClientDao) AssignClientToClinician(clientID int, clinicianID int) (int, error) {
	var query = "INSERT INTO cfc.client_has_clinician(client_clientid, clinician_clinicianid) VALUES($1, $2)"
	var parameters = []interface{}{clientID, clinicianID}
	return cd.db.Insert(query, parameters)
}

func (cd *ClientDao) GetClientNameByClientID(clientID int) (*Model.Person, error) {
	var query = "SELECT userid, firstname, lastname FROM cfc.\"ClientNames\" WHERE clientid = $1"
	var parameters = []interface{}{clientID}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return new(Model.Person), err
	}

	person := Model.NewPerson("", "", result[0][1], result[0][2], "", "", "", "", "", "")
	uid, _ := strconv.ParseInt(result[0][0], 10, 64)
	person.SetUserID(int(uid))

	return person, nil
}

func (cd *ClientDao) GetAllClientNames() ([]*Model.Person, error) {
	var clientNamesList []*Model.Person
	var query = "SELECT userid, firstname, lastname FROM cfc.\"ClientNames\""
	var parameters []interface{}

	result, err := cd.db.Select(query, parameters)
	if err != nil {
		return clientNamesList, err
	}

	for _, res := range result {
		person := Model.NewPerson("", "", res[1], res[2], "", "", "", "", "", "")
		uid, _ := strconv.ParseInt(res[0], 10, 64)
		person.SetUserID(int(uid))
		clientNamesList = append(clientNamesList, person)
	}

	return clientNamesList, nil
}
