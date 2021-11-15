package handlers

// import (
// 	"CFC/backend/CFC/backend/DB"
// 	Facade "CFC/backend/CFC/backend/facade"
// 	"encoding/json"
// 	"net/http"
// )


// type SafetyPlanHandler struct {
// 	Database DB.DatabaseConnection
// }

// func (sph *SafetyPlanHandler) ClientGetSafetyPlan(w http.ResponseWriter, r *http.Request) {
// 	claims, er := Auth.IsAuthorized(w, r)
// 	if !er {
// 		return
// 	}

// 	client := Facade.NewClientFacade(sph.Database)
// 	var userID int = int(claims["userID"].(float64))
// 	pers, err := person.
// }
