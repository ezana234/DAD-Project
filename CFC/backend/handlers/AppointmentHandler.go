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

type AppointmentHandler struct {
	Database DB.DatabaseConnection
}

func NewAppointmentHandler(db DB.DatabaseConnection) *AppointmentHandler {
	return &AppointmentHandler{Database: db}
}

func (ah *AppointmentHandler) GetAppointment(w http.ResponseWriter, r *http.Request) {
	_, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}
	appointmentID := r.URL.Query().Get("appointmentID")
	intAppointmentID, err := strconv.Atoi(appointmentID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusUnauthorized)
	}

	applicationFacade := Facade.NewAppointmentFacade(ah.Database)

	app, intReturn := applicationFacade.GetAppointmentByID(intAppointmentID)
	if intReturn == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	b, erro := json.Marshal(app)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (ah *AppointmentHandler) ClientGetAppointments(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	var userID int = int(claims["userID"].(float64))
	var role = fmt.Sprintf("%v", claims["role"])
	pf := Facade.NewPersonFacade(ah.Database)
	appointments, erro := pf.GetAppointmentsByUserID(userID, role)
	if erro == 0 {
		return
	}

	b, err := json.Marshal(appointments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (ah *AppointmentHandler) AddAppointment(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	var role = fmt.Sprintf("%v", claims["role"])
	if role == "2" {
		var app Model.Appointment

		body := json.NewDecoder(r.Body).Decode(&app)
		println(app.AppointmentMedium)
		if body != nil {
			http.Error(w, body.Error(), http.StatusBadRequest)
			return
		}

		af := Facade.NewAppointmentFacade(ah.Database)
		intReturn := af.AddAppointment(app)
		if intReturn == 0 {
			http.Error(w, "error when adding appointment", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (ah *AppointmentHandler) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	var role = fmt.Sprintf("%v", claims["role"])
	if role == "2" {
		var app Model.Appointment

		body := json.NewDecoder(r.Body).Decode(&app)
		if body != nil {
			http.Error(w, body.Error(), http.StatusBadRequest)
			return
		}
		af := Facade.NewAppointmentFacade(ah.Database)
		intReturn := af.UpdateAppointment(app.GetAppointmentID(), app)
		if intReturn != 1 {
			http.Error(w, "error when updating appointment", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (ah *AppointmentHandler) DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	type AppointmentToDelete struct {
		AppointmentID int
	}

	var role = fmt.Sprintf("%v", claims["role"])
	if role == "2" {
		var appToDelete AppointmentToDelete
		body := json.NewDecoder(r.Body).Decode(&appToDelete)
		if body != nil {
			http.Error(w, body.Error(), http.StatusBadRequest)
			return
		}

		af := Facade.NewAppointmentFacade(ah.Database)
		intReturn := af.DeleteAppointment(appToDelete.AppointmentID)
		if intReturn != 1 {
			http.Error(w, "error when deleting appointment", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (ah *AppointmentHandler) ClinicianGetAllAppointments(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if !er {
		return
	}

	var role = fmt.Sprintf("%v", claims["role"])
	if role == "2" {
		applicationFacade := Facade.NewAppointmentFacade(ah.Database)

		apps, intReturn := applicationFacade.GetAllAppointments()
		if intReturn == 0 {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		b, erro := json.Marshal(apps)
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
