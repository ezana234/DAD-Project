package handlers

import (
	Facade "CFC/backend/CFC/backend/facade"
	"encoding/json"
	"net/http"
)

func (db Database) client(w http.ResponseWriter, r *http.Request) {
	claims, er := isAuthorized(w, r)
	if er == false {
		return
	}
	// body := json.NewDecoder(r.Body).Decode(&clientStruct)
	// if body != nil {
	// 	http.Error(w, body.Error(), http.StatusBadRequest)
	// 	return
	// }
	person := Facade.NewPersonFacade(db.database)
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
