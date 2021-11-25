package handlers

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	Facade "CFC/backend/CFC/backend/facade"
	"encoding/json"
	"net/http"
)

type ClientHandler struct {
	Database DB.DatabaseConnection
}

// This function gets the client from the JWT
func (ch *ClientHandler) GetClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	claims, er := Auth.IsAuthorized(w, r)
	if er == false {
		return
	}
	person := Facade.NewPersonFacade(ch.Database)
	var userID int = int(claims["userID"].(float64))
	pers, err := person.GetPerson(userID)
	if err == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
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

	return
}
