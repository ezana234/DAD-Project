package handlers

import (
	"CFC/backend/CFC/backend/DB"
	Facade "CFC/backend/CFC/backend/facade"
	Auth "CFC/backend/CFC/backend/auth"
	Model "CFC/backend/CFC/backend/model"
	"encoding/json"
	"net/http"
)


type AuthHandler struct {
	Database DB.DatabaseConnection
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	type Login struct {
		Email    string
		Password string
	}
	var logStruct Login
	body := json.NewDecoder(r.Body).Decode(&logStruct)
	if body != nil {
		http.Error(w, body.Error(), http.StatusBadRequest)
		return
	}
	person := Facade.NewPersonFacade(ah.Database)
	pers := person.GetPersonByEmail(logStruct.Email, logStruct.Password)
	if pers.GetUserID() == 0 {
		http.Error(w, "Bad Login", http.StatusUnauthorized)
		return
	} else {
		tokenString, erro := Auth.GenerateJWT(pers.GetUserID(), pers.GetEmail(), pers.GetRole())
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
		resp := make(map[string]string)
		resp["token"] = tokenString
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

// This returns a jwt upon a successful login
func (ah *AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
	type sign struct {
		Username    string
		FirstName   string
		LastName    string
		Email       string
		Address     string
		Password    string
		PhoneNumber string
		DOB         string
	}
	var signStruct sign
	body := json.NewDecoder(r.Body).Decode(&signStruct)
	if body != nil {
		http.Error(w, body.Error(), http.StatusBadRequest)
		return
	}

	person := Facade.NewPersonFacade(ah.Database)
	newPers := Model.NewPerson(
		signStruct.Username,
		signStruct.Password,
		signStruct.FirstName,
		signStruct.LastName,
		signStruct.Email,
		signStruct.Address,
		signStruct.PhoneNumber,
		"1",
		" ",
		signStruct.DOB)
	userID, err := person.CreateNewPerson(*newPers)
	if err != 1 {
		http.Error(w, "Couldn't Create Person", http.StatusBadRequest)
		return
	} else {
		// fmt.Println(userID)
		client := Facade.NewClientFacade(ah.Database)
		clientModel := Model.NewClient(userID)
		client.AddClient(*clientModel)
		resp := make(map[string]string)
		resp["message"] = "Client added to Database"
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}