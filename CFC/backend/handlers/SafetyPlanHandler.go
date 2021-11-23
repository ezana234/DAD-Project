package handlers

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	Facade "CFC/backend/CFC/backend/facade"
	Model "CFC/backend/CFC/backend/model"
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

func (sph *SafetyPlanHandler) ClinicianAddSafetyPlan(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if er == false {
		return
	}

	type SafetyPlanIN struct {
		Triggers             string
		WarningSigns         string
		DestructiveBehaviors string
		InternalStrategies   string
		ClientID             int
	}

	var role = fmt.Sprintf("%v", claims["role"])
	if role == "2" {
		var spStruct SafetyPlanIN
		body := json.NewDecoder(r.Body).Decode(&spStruct)
		if body != nil {
			http.Error(w, body.Error(), http.StatusBadRequest)
			return
		}

		userID, _ := strconv.ParseInt(fmt.Sprintf("%v", claims["userID"]), 10, 64)
		personFacade := Facade.NewPersonFacade(sph.Database)
		clinician, returnInt := personFacade.GetClinicianByUserID(int(userID))
		if returnInt != 1 {
			http.Error(w, "error when selecting updated clinician", http.StatusInternalServerError)
			return
		}

		spf := Facade.NewSafetyPlanFacade(sph.Database)
		sp := Model.NewSafetyPlan(spStruct.Triggers, spStruct.WarningSigns, spStruct.DestructiveBehaviors, spStruct.InternalStrategies, "", clinician.GetClinicianID(), spStruct.ClientID, clinician.GetClinicianID())
		intReturn := spf.AddSafetyPlan(sp)
		if intReturn != 1 {
			http.Error(w, "error when adding safety plan", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (sph *SafetyPlanHandler) ClinicianUpdateSafetyPlan(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if er == false {
		return
	}

	type SafetyPlanIN struct {
		SafetyID             int
		Triggers             string
		WarningSigns         string
		DestructiveBehaviors string
		InternalStrategies   string
		ClientID             int
		ClinicianID          int
	}
	var role = fmt.Sprintf("%v", claims["role"])
	if role == "2" {
		var spStruct SafetyPlanIN
		body := json.NewDecoder(r.Body).Decode(&spStruct)
		if body != nil {
			http.Error(w, body.Error(), http.StatusBadRequest)
			return
		}
		userID, _ := strconv.ParseInt(fmt.Sprintf("%v", claims["userID"]), 10, 64)
		personFacade := Facade.NewPersonFacade(sph.Database)
		clinician, returnInt := personFacade.GetClinicianByUserID(int(userID))
		if returnInt != 1 {
			http.Error(w, "error when selecting updated clinician", http.StatusInternalServerError)
			return
		}
		spf := Facade.NewSafetyPlanFacade(sph.Database)
		sp := Model.NewSafetyPlan(spStruct.Triggers, spStruct.WarningSigns, spStruct.DestructiveBehaviors, spStruct.InternalStrategies, "", clinician.GetClinicianID(), spStruct.ClientID, spStruct.ClinicianID)
		intReturn := spf.UpdateSafetyPlan(spStruct.SafetyID, sp)
		if intReturn != 1 {
			http.Error(w, "error when updating safety plan", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (sph *SafetyPlanHandler) ClinicianDeleteSafetyPlan(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if er == false {
		return
	}

	safetyID := r.URL.Query().Get("safetyID")
	intSafetyID, err := strconv.Atoi(safetyID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusUnauthorized)
	}
	var role string = fmt.Sprintf("%v", claims["role"])
	if role == "2" {
		spf := Facade.NewSafetyPlanFacade(sph.Database)
		intReturn := spf.DeleteSafetyPlan(intSafetyID)
		if intReturn != 1 {
			http.Error(w, "error when updating safety plan", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}
