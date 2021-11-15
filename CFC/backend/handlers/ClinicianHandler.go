package handlers

// import (
// 	"CFC/backend/CFC/backend/DB"
// 	Auth "CFC/backend/CFC/backend/auth"
// 	Facade "CFC/backend/CFC/backend/facade"
// 	Model "CFC/backend/CFC/backend/model"
// 	"encoding/json"
// 	"net/http"
// ) 


// type ClinicianHandler struct {
// 	Database DB.DatabaseConnection
// }


// func (ch *ClinicianHandler) GetClients(w http.ResponseWriter, r *http.Request) {
// 	_, er := Auth.IsAuthorized(w, r)
// 	if !er {
// 		return
// 	}

// 	cf := Facade.NewClientFacade(ch.Database)
// 	cList, _ := cf.GetAllUserClients()

