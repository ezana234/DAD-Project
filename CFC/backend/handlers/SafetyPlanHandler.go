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

type SafetyPlanHandler struct {
	Database DB.DatabaseConnection
}

func (sph *SafetyPlanHandler) ClientGetSafetyPlan(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	var userID int = int(claims["userID"].(float64))
	sf := Facade.NewPersonFacade(sph.Database)
	safetyplan, _ := sf.GetSafetyPlansByUserID(userID, 1)

	b, err := json.Marshal(safetyplan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (sph *SafetyPlanHandler) ClinicianGetSafetyPlans(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if er == false {
		return
	}

	// Check if the person is a clinician
	var role = fmt.Sprintf("%v", claims["role"])
	println(role)
	if role == "2" {
		println("true")
		spf := Facade.NewSafetyPlanFacade(sph.Database)
		spList, _ := spf.GetAllSafetyPlans()
		println(len(spList))
		for _, a := range spList {
			println(a.SafetyPlanToString())
		}
		b, err := json.Marshal(spList)
		println(string(b))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (sph *SafetyPlanHandler) ClinicianGetSafetyPlan(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if er == false {
		return
	}
	userID := r.URL.Query().Get("userID")
	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusUnauthorized)
	}
	// Check if the person is a client
	var role string = fmt.Sprintf("%v", claims["role"])
	//intRole, err := strconv.Atoi(role)
	if role == "2" {
		spf := Facade.NewSafetyPlanFacade(sph.Database)
		safetyPlan, _ := spf.GetSafetyPlanByUserID(intUserID)
		//person := Facade.NewPersonFacade(sph.Database)
		//safetyPlan, _ := person.GetSafetyPlansByUserID(intUserID, intRole)
		b, erro := json.Marshal(safetyPlan)
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
