package handlers

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	Facade "CFC/backend/CFC/backend/facade"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ClientHandler struct {
	Database DB.DBConnection
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

func (ch *ClientHandler) GetClientName(w http.ResponseWriter, r *http.Request) {
	_, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	type ClientName struct {
		UserID    int
		FirstName string
		LastName  string
	}

	clientID := r.URL.Query().Get("clientID")
	intClientID, err := strconv.Atoi(clientID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusUnauthorized)
	}

	cf := Facade.NewClientFacade(ch.Database)
	client, intReturn := cf.GetUserByClientID(intClientID)
	if intReturn != 1 {
		http.Error(w, "error when selecting client name", http.StatusInternalServerError)
		return
	}

	b, erro := json.Marshal(ClientName{UserID: client.GetUserID(), FirstName: client.GetFirstName(), LastName: client.GetLastName()})
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (ch *ClientHandler) GetAllClientNames(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	type ClientName struct {
		UserID    int
		FirstName string
		LastName  string
	}

	var role = fmt.Sprintf("%v", claims["role"])
	if role == "2" {
		var clientNameList []ClientName

		cf := Facade.NewClientFacade(ch.Database)
		clients, err := cf.GetAllClientNames()
		if err == 0 {
			return
		}

		for _, a := range clients {
			clientNameList = append(clientNameList, ClientName{UserID: a.GetUserID(), FirstName: a.GetFirstName(), LastName: a.GetLastName()})
		}

		b, er := json.Marshal(clientNameList)
		if er != nil {
			http.Error(w, er.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (ch *ClientHandler) GetClientNameByUserID(w http.ResponseWriter, r *http.Request) {
	_, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	type ClientName struct {
		UserID    int
		FirstName string
		LastName  string
	}

	userID := r.URL.Query().Get("userID")
	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusUnauthorized)
	}

	pf := Facade.NewPersonFacade(ch.Database)
	client, erro := pf.GetClientNameByUserID(intUserID)
	if erro == 0 {
		http.Error(w, "error when selecting client name", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(ClientName{UserID: client.GetUserID(), FirstName: client.GetFirstName(), LastName: client.GetLastName()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (ch *ClientHandler) GetClientInfoByUserID(w http.ResponseWriter, r *http.Request) {
	_, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	userID := r.URL.Query().Get("userID")
	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusUnauthorized)
	}

	pf := Facade.NewPersonFacade(ch.Database)
	client, erro := pf.GetClientByUserID(intUserID)
	if erro == 0 {
		http.Error(w, "error when selecting client", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}
