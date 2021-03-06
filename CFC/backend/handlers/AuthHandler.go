package handlers

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	Facade "CFC/backend/CFC/backend/facade"
	Model "CFC/backend/CFC/backend/model"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	Database DB.DBConnection
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
	pers, _ := person.LoginPersonByEmail(logStruct.Email, logStruct.Password)
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
		_, err = w.Write(b)
		if err != nil {
			return
		}
	}
}

//// This returns a jwt upon a successful login
//func (ah *AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
//	type sign struct {
//		Username    string
//		FirstName   string
//		LastName    string
//		Email       string
//		Address     string
//		Password    string
//		PhoneNumber string
//		DOB         string
//	}
//	var signStruct sign
//	body := json.NewDecoder(r.Body).Decode(&signStruct)
//	if body != nil {
//		http.Error(w, body.Error(), http.StatusBadRequest)
//		return
//	}
//
//	person := Facade.NewPersonFacade(ah.Database)
//	newPers := Model.NewPerson(
//		signStruct.Username,
//		signStruct.Password,
//		signStruct.FirstName,
//		signStruct.LastName,
//		signStruct.Email,
//		signStruct.Address,
//		signStruct.PhoneNumber,
//		"1",
//		" ",
//		signStruct.DOB)
//	userID, err := person.CreateNewPerson(*newPers)
//	if err != 1 {
//		http.Error(w, "Couldn't Create Person", http.StatusBadRequest)
//		return
//	} else {
//		client := Facade.NewClientFacade(ah.Database)
//		clientModel := Model.NewClient(userID)
//		client.AddClient(*clientModel)
//		resp := make(map[string]string)
//		resp["message"] = "Client added to Database"
//		b, err := json.Marshal(resp)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(b)
//	}
//}

func (ah *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	type signUp struct {
		Username    string
		FirstName   string
		LastName    string
		Email       string
		Address     string
		Password    string
		PhoneNumber string
		DOB         string
		Referral    string
	}

	var signStruct signUp
	body := json.NewDecoder(r.Body).Decode(&signStruct)
	if body != nil {
		http.Error(w, body.Error(), http.StatusBadRequest)
		return
	}

	clinicianFacade := Facade.NewClinicianFacade(ah.Database)
	clinicianID, _ := clinicianFacade.GetClinicianIDByReferral(signStruct.Referral)
	if clinicianID == 0 {
		http.Error(w, "Couldn't Create Person: No Clinician With Given Referral Code Found", http.StatusBadRequest)
		return
	}

	personFacade := *Facade.NewPersonFacade(ah.Database)
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
	userID, returnInt := personFacade.CreateNewPerson(*newPers)
	if returnInt == -1 {
		http.Error(w, "Couldn't Create Person: Username Already Taken", http.StatusBadRequest)
		return
	} else if returnInt == 0 {
		http.Error(w, "Couldn't Create Person: bad request", http.StatusBadRequest)
		return
	} else {
		client := Facade.NewClientFacade(ah.Database)
		clientModel := Model.NewClient(userID)
		clientID, intReturn := client.AddClient(*clientModel)
		if intReturn != 1 {
			http.Error(w, "Error when adding client", http.StatusInternalServerError)
			return
		}
		intReturn = client.AssignClinicianToClient(clientID, clinicianID)
		if intReturn != 1 {
			http.Error(w, "Error when assigning clinician", http.StatusInternalServerError)
			return
		}

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
