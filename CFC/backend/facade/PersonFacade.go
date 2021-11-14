package facade

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
	"encoding/json"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type PersonFacade struct {
	personDao   DAO.PersonDao
	authManager *Auth.AuthenticationManager
}

func NewPersonFacade(db DB.DatabaseConnection) *PersonFacade {
	return &PersonFacade{personDao: *DAO.NewPersonDao(db)}
}

//// GetAuthManager
//// TEMPORARY FUNCTION
//func (pf *PersonFacade) GetAuthManager() *Auth.AuthenticationManager {
//	return pf.authManager
//}

func (pf *PersonFacade) GetPerson(userID int) (*Model.Person, int) {
	p, err := pf.personDao.GetUserByID(userID)
	if err != nil {
		log.Printf("Error: %s when getting person\n", err)
		return new(Model.Person), 0
	}
	//p.SetPassword("null")

	return p, 1
}

func (pf *PersonFacade) GetAllPersons() ([]*Model.Person, int) {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() {
		var pList []*Model.Person

		tmp, err := pf.personDao.GetAll()
		if err != nil {
			log.Printf("Error: %s when getting all persons", err)
			return pList, 0
		}

		for _, res := range tmp {
			res.SetPassword("null")
			pList = append(pList, res)
		}

		return pList, 1
	}

	log.Printf("Error: user is not authorized to get persons")
	return []*Model.Person{}, -1

}

func (pf *PersonFacade) GetNPersons(num int) ([]*Model.Person, int) {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() {
		var pList []*Model.Person

		tmp, err := pf.personDao.GetAll()
		if err != nil {
			log.Printf("Error: %s when getting number of persons", err)
			return pList, 0
		}

		for _, res := range tmp[:num] {
			res.SetPassword("null")
			pList = append(pList, res)
		}

		return pList, 1
	}

	log.Printf("Error: user is not authorized to get persons")
	return []*Model.Person{}, -1
}

func (pf *PersonFacade) GetPersonByEmail(email string, password string) *Model.Person {
	p, err := pf.personDao.GetPersonByEmail(email)
	if err != nil {
		log.Printf("Error: %s when getting person by email", err)
		return new(Model.Person)
	}

	//p.SetPassword("null")

	return p

	// log.Printf("Error: User is not authorized to get person by email")
	// return new(Model.Person), -1
}

func (pf *PersonFacade) AddPerson(p Model.Person) int {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() {
		p.SetUserID(pf.personDao.GetNextUserID())

		err := pf.personDao.Add(p)
		if err != nil {
			log.Printf("Error: %s when adding person")
			return 0
		}

		return 1
	}

	log.Printf("Error: User is not authorized to add person")
	return -1
}

func (pf *PersonFacade) UpdatePerson(userID int, p Model.Person) int {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() || pf.authManager.IsCurrentUser(userID) {
		pOld, err := pf.personDao.GetUserByID(userID)
		if err != nil {
			log.Printf("Error: %s when getting person")
			return 0
		}
		var pNew = Model.NewPerson(p.GetUserName(), pOld.GetPassword(), p.GetFirstName(), p.GetLastName(), p.GetEmail(), p.GetAddress(), p.GetPhoneNumber(), p.GetRole(), p.GetExpiration(), p.GetDOB())

		err = pf.personDao.Update(userID, pNew)
		if err != nil {
			log.Printf("Error: %s when updating person", err)
			return 0
		}

		return 1
	}

	log.Printf("Error: User is not authorized to update person")
	return -1
}

// DeletePerson
// function will delete a person from the database
// returns -1 if user is not authorized to delete
// returns 0 if deletion failed
// returns 1 if deletion was successful
func (pf *PersonFacade) DeletePerson(userID int) int {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() {
		err := pf.personDao.Delete(userID)
		if err != nil {
			log.Printf("Error: %s when deleting person", err)
			return 0
		}

		return 1
	}

	log.Printf("Error: User is not authorized to delete person")
	return -1
}

// CreateNewPerson
// this functions adds a new user to the db when they create their account for the first time
// returns 0 if creation was unsuccessful, 1 if it was successful
func (pf *PersonFacade) CreateNewPerson(p Model.Person) int {
	usernameIsPresent, err := pf.personDao.UsernameExists(p.UserName)
	println(usernameIsPresent)
	if err != nil {
		log.Printf("Error: %s when creating new person", err)
		return 0
	}

	if !usernameIsPresent {
		return -1
	}

	p.SetUserID(pf.personDao.GetNextUserID())
	p.SetPassword(HashPassword(p.GetPassword()))

	err = pf.personDao.Add(p)
	if err != nil {
		log.Printf("Error: %s when creating new person", err)
		return 0
	}

	return 1
}

// LoginPersonByUserName
// this function will query all persons with a matching username and then check if the passwords match.
// if there are no persons that have the desired username, then this function will return 0.
// if the password does not match, this function will return 0.
// if there is a password match, this function will return 1.
// TODO Check if password has expired and if so, prompt user to reset password
func (pf *PersonFacade) LoginPersonByUserName(userName string, password string) int {
	p, err := pf.personDao.GetPersonByUserName(userName)
	if err != nil {
		log.Printf("Error: %s when logging in by username", err)
		return 0
	}

	if CheckPasswords(p.GetPassword(), password) {
		if !IsExpired(p.GetExpiration()) {
			return -1
		}

		pf.authManager.LoginUser(p)
		return 1
	}

	return 0
}

func (pf *PersonFacade) LoginPersonByEmail(email string, password string) (*Model.Person, int){
	p, err := pf.personDao.GetPersonByEmail(email)
	if err != nil {
		log.Printf("Error: %s when logging in by email", err)
		return new(Model.Person), 0
	}

	if CheckPasswords(p.GetPassword(), password) {
		if !IsExpired(p.GetExpiration()) {
			return new(Model.Person), -1
		}

		return p, 1
	}
	
	return new(Model.Person), 0
}

func (pf PersonFacade) UpdatePassword(password string) int {
	p := pf.authManager.GetCurrentUser()
	p.SetPassword(HashPassword(password))

	err := pf.personDao.Update(p.GetUserID(), p)
	if err != nil {
		log.Printf("Error: %s when updating password", err)
		return 0
	}

	return 1
}

func (pf PersonFacade) ResetPassword(username string) int {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() || pf.authManager.GetCurrentUser().GetUserName() == username {
		p, err := pf.personDao.GetPersonByUserName(username)
		if err != nil {
			log.Printf("Error: %s when getting persons by email", err)
			return 0
		}

		p.SetPassword("temp")

		err = pf.personDao.Update(p.GetUserID(), p)
		if err != nil {
			log.Printf("Error: %s whem updating person", err)
			return 0
		}

		return 1
	}

	log.Printf("Error: user is not authorized to reset password")
	return -1
}

func HashPassword(password string) string {
	bPassword := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bPassword, 14)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func CheckPasswords(hp, cp string) bool {
	bhp := []byte(hp)
	bcp := []byte(cp)

	err := bcrypt.CompareHashAndPassword(bhp, bcp)
	if err != nil {
		return false
	}

	return true
}

func IsExpired(expiration string) bool {
	timeNow := time.Now()
	tExpire, _ := time.Parse(time.RFC3339, expiration)
	if timeNow.After(tExpire) {
		return false
	}

	return true
}

func PersonFromJSON(pJson []byte) Model.Person {
	var p = Model.Person{}

	var err = json.Unmarshal(pJson, &p)
	if err != nil {
		panic(err)
	}

	return p
}

func PersonArrayFromJSON(pJson []byte) []Model.Person {
	var pList []Model.Person

	var err = json.Unmarshal(pJson, &pList)
	if err != nil {
		panic(err)
	}

	return pList
}

func PersonToJSON(obj interface{}) []byte {
	pJson, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return pJson
}
