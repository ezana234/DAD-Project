package facade

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	DAO "CFC/backend/CFC/backend/dao"
	Model "CFC/backend/CFC/backend/model"
	"golang.org/x/crypto/bcrypt"
)

type PersonFacade struct {
	personDao   DAO.PersonDao
	authManager *Auth.AuthenticationManager
}

func NewPersonFacade(db DB.DatabaseConnection, authManager *Auth.AuthenticationManager) *PersonFacade {
	return &PersonFacade{personDao: *DAO.NewPersonDao(db), authManager: authManager}
}

func (pf *PersonFacade) GetAuthManager() *Auth.AuthenticationManager {
	return pf.authManager
}

func (pf *PersonFacade) GetPerson(userID int) (*Model.Person, int) {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() || pf.authManager.IsCurrentUser(userID) {
		p := pf.personDao.GetByID(userID)
		p.SetPassword("null")

		return p, 1
	}

	return new(Model.Person), -1
}

func (pf *PersonFacade) GetPersons() ([]*Model.Person, int) {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() {
		var pList []*Model.Person

		tmp := pf.personDao.GetAll()
		for _, res := range tmp {
			res.SetPassword("null")
			pList = append(pList, res)
		}

		return pList, 1
	}

	return []*Model.Person{}, -1
}

func (pf *PersonFacade) GetPersonByEmail(email string) (*Model.Person, int) {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() || pf.authManager.GetCurrentUser().GetEmail() == email {
		pList := pf.personDao.GetPersonsByEmail(email)
		if len(pList) == 0 {
			return new(Model.Person), 0
		}

		return pList[0], 1
	}

	return new(Model.Person), -1
}

func (pf *PersonFacade) AddPerson(p Model.Person) int {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() {
		p.SetUserID(pf.personDao.GetNextUserID())

		err := pf.personDao.Add(p)
		if err != nil {
			return 0
		}

		return 1
	}

	return -1
}

func (pf *PersonFacade) UpdatePerson(userID int, p Model.Person) int {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() || pf.authManager.IsCurrentUser(userID) {
		var pOld = pf.personDao.GetByID(userID)
		var pNew = Model.NewPerson(p.GetUserName(), pOld.GetPassword(), p.GetFirstName(), p.GetLastName(), p.GetEmail(), p.GetAddress(), p.GetPhoneNumber(), p.GetRole())

		err := pf.personDao.Update(userID, pNew)
		if err != nil {
			return 0
		}

		return 1
	}

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

			return 0
		}

		return 1
	}

	return -1
}

// CreateNewPerson
// this functions adds a new user to the db when they create their account for the first time
// returns 0 if creation was unsuccessful, 1 if it was successful
func (pf *PersonFacade) CreateNewPerson(p Model.Person) int {
	p.SetUserID(pf.personDao.GetNextUserID())
	p.SetPassword(HashPassword(p.GetPassword()))

	err := pf.personDao.Add(p)
	if err != nil {
		return 0
	}

	return 1
}

// LoginPersonByUserName
// this function will query all persons with a matching username and then check if the passwords match.
// if there are no persons that have the desired username, then this function will return 0.
// if there are persons with the desired username, but the password does not match, then this function will return -1.
// if there are persons with the desired username, and the password matches, then this function will return 1.
// TODO Check if password has expired and if so, prompt user to reset password
func (pf *PersonFacade) LoginPersonByUserName(userName string, password string) int {
	pList := pf.personDao.GetPersonsByUserName(userName)
	if len(pList) == 0 {
		return 0
	}

	for _, p := range pList {
		if p.GetPassword() == "temp" {
			return -3
		}
		if CheckPasswords(p.GetPassword(), password) {
			pf.authManager.LoginUser(p)

			return 1
		}
	}

	return -1
}

// LoginPersonByEmail
// TODO Check if password has expired and if so, prompt user to reset password
func (pf *PersonFacade) LoginPersonByEmail(email string, password string) int {
	pList := pf.personDao.GetPersonsByEmail(email)
	if len(pList) == 0 {
		return 0
	}

	for _, p := range pList {
		if CheckPasswords(p.GetPassword(), password) {
			pf.authManager.LoginUser(p)

			return 1
		}
	}

	return -1
}

func (pf PersonFacade) UpdatePassword(password string) int {
	p := pf.authManager.GetCurrentUser()
	p.SetPassword(HashPassword(password))

	err := pf.personDao.Update(p.GetUserID(), p)
	if err != nil {
		return 0
	}

	return 1
}

func (pf PersonFacade) ResetPassword(email string) int {
	if pf.authManager.IsCurrentUserAdmin() || pf.authManager.IsCurrentUserClinician() || pf.authManager.GetCurrentUser().GetEmail() == email {
		p := pf.personDao.GetPersonsByEmail(email)[0]
		p.SetPassword("temp")

		err := pf.personDao.Update(p.GetUserID(), p)
		if err != nil {
			return 0
		}

		return 1
	}

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
