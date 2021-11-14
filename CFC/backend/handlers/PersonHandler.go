package handlers

import (
	DB "CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	Facade "CFC/backend/CFC/backend/facade"
	"encoding/json"
	"net/http"
)

type PersonHandler struct {
	Database DB.DatabaseConnection
}


func NewPersonHandler(db DB.DatabaseConnection) *PersonHandler {
	return &PersonHandler{Database: db}
}


func (ph *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}
	// body := json.NewDecoder(r.Body).Decode(&clientStruct)
	// if body != nil {
	// 	http.Error(w, body.Error(), http.StatusBadRequest)
	// 	return
	// }
	person := Facade.NewPersonFacade(ph.Database)
	var i int = int(claims["userID"].(float64))
	pers, err := person.GetPerson(i)
	if err == 0 {
		http.Error(w, pers.Error(), http.StatusNotFound)
		return
	}
	type PersonMessage struct {
		UserID      int
		UserName    string
		FirstName   string
		LastName    string
		Email       string
		Address     string
		PhoneNumber string
		Role        string
	}
	persJson := PersonMessage{
		UserID:      pers.GetUserID(),
		UserName:    pers.GetUserName(),
		FirstName:   pers.GetFirstName(),
		LastName:    pers.GetLastName(),
		Email:       pers.GetEmail(),
		Address:     pers.GetAddress(),
		PhoneNumber: pers.GetPhoneNumber(),
		Role:        pers.GetRole()}

	b, erro := json.Marshal(persJson)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
