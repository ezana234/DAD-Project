package handlers

import (
	"CFC/backend/CFC/backend/DB"
	Auth "CFC/backend/CFC/backend/auth"
	Facade "CFC/backend/CFC/backend/facade"
	"encoding/json"
	"net/http"
)


type SafetyPlanHandler struct {
	Database DB.DatabaseConnection
}

func (sph *SafetyPlanHandler) ClientGetSafetyPlan(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	sf := Facade.NewPersonFacade(sph.Database)
	var userID int = int(claims["userID"].(float64))
	safetyplan, _ := sf.GetSafetyPlansByUserID(userID, 1)

	b, err := json.Marshal(safetyplan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}