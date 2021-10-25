package main

import (
	"CFC/backend/CFC/backend/DB"
	Facade "CFC/backend/CFC/backend/facade"
	Model "CFC/backend/CFC/backend/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Database struct {
	database DB.DatabaseConnection
}

func main() {
	db := *DB.NewDatabaseConnection("sql5446146", "WUi5dvp7gj", "sql5.freemysqlhosting.net:3306", "sql5446146")
	cf := *Facade.NewClinicianFacade(db)
	newClinician := *Model.NewClinician(1002)
	cf.AddClinician(newClinician)
	mux := http.NewServeMux()
	dbHandler := &Database{database: db}
	// Routes
	mux.HandleFunc("/login", dbHandler.login)

	log.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func (db *Database) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	person := Facade.NewPersonFacade(db.database)
	pers := person.GetPersonByEmail(logStruct.Email, logStruct.Password)
	if pers.UserID() == 0 {
		http.Error(w, "Bad Login", http.StatusUnauthorized)
		return
	} else {
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
			UserID:      pers.UserID(),
			UserName:    pers.UserName(),
			FirstName:   pers.FirstName(),
			LastName:    pers.LastName(),
			Email:       pers.Email(),
			Address:     pers.Address(),
			PhoneNumber: pers.PhoneNumber(),
			Role:        pers.Role()}
		b, err := json.Marshal(persJson)
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println(pers.Role())
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}
