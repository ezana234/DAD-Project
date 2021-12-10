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

type ClinicianHandler struct {
	Database DB.DBConnection
}

func (ch *ClinicianHandler) GetClinicians(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	var role = fmt.Sprintf("%v", claims["role"])
	if role == "2" {
		clinicianFacade := Facade.NewClinicianFacade(ch.Database)

		clinicians, intReturn := clinicianFacade.GetAllClinicians()
		if intReturn == 0 {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		b, erro := json.Marshal(clinicians)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (ch *ClinicianHandler) GetClinicianName(w http.ResponseWriter, r *http.Request) {
	_, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	type ClinicianName struct {
		UserID    int
		FirstName string
		LastName  string
	}

	clinicianID := r.URL.Query().Get("clinicianID")
	intClinicianID, err := strconv.Atoi(clinicianID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusUnauthorized)
	}

	cf := Facade.NewClinicianFacade(ch.Database)
	clinician, intReturn := cf.GetClinicianNameByClinicianID(intClinicianID)
	if intReturn != 1 {
		http.Error(w, "error when selecting clinician name", http.StatusInternalServerError)
		return
	}

	b, erro := json.Marshal(ClinicianName{UserID: clinician.GetUserID(), FirstName: clinician.GetFirstName(), LastName: clinician.GetLastName()})
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (ch *ClinicianHandler) GetClinicianNameByUserID(w http.ResponseWriter, r *http.Request) {
	_, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	type ClinicianName struct {
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
	clinicianName, erro := pf.GetClinicianNameByUserID(intUserID)
	if erro != 1 {
		http.Error(w, "error when selecting clinician name", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(ClinicianName{UserID: clinicianName.GetUserID(), FirstName: clinicianName.GetFirstName(), LastName: clinicianName.GetLastName()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (ch *ClinicianHandler) GetClinicianNames(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	type ClinicianName struct {
		UserID    int
		FirstName string
		LastName  string
	}

	var role = fmt.Sprintf("%v", claims["role"])
	if role == "2" {
		var clinicianNameList []ClinicianName
		cf := Facade.NewClinicianFacade(ch.Database)

		clinicianNames, returnInt := cf.GetAllClinicianNames()
		if returnInt == 0 {
			return
		}

		for _, a := range clinicianNames {
			clinicianNameList = append(clinicianNameList, ClinicianName{UserID: a.GetUserID(), FirstName: a.GetFirstName(), LastName: a.GetLastName()})
		}

		b, er := json.Marshal(clinicianNameList)
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

func (ch *ClinicianHandler) GetClinicianInfoByUserID(w http.ResponseWriter, r *http.Request) {
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
	clinician, erro := pf.GetClinicianByUserID(intUserID)
	if erro == 0 {
		http.Error(w, "error when selecting clinician", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(clinician)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
