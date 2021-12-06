package main

import (
	"CFC/backend/CFC/backend/DB"
	// "strings"
	// "time"

	// 	Auth "CFC/backend/CFC/backend/auth"
	// 	Facade "CFC/backend/CFC/backend/facade"
	// 	Handlers "CFC/backend/CFC/backend/handlers"

	Auth "CFC/backend/CFC/backend/auth"
	Facade "CFC/backend/CFC/backend/facade"
	Handlers "CFC/backend/CFC/backend/handlers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Database struct {
	database DB.DatabaseConnection
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	db := *DB.NewDatabaseConnection("ydmscaoenbipqz", "f9ac329ae1c957bdd5015e4f91bb7968850dd6eb2773105ff6f2b4efb036de47", "ec2-52-54-237-144.compute-1.amazonaws.com", "5432", "d85fspl6bklvdv")
	mux := mux.NewRouter()
	dbHandler := &Database{database: db}
	mux.Use(accessControlMiddleware)
	//pf := Facade.NewPersonFacade(db)
	//pNew := *model.NewPerson("clinicianuser2", "c2password", "Clinician2", "User2", "clin2@gmail.com", "123 Street", "123456789", "2", "", "04/03/2002")
	//pf.CreateNewPerson(pNew)

	mux.HandleFunc("/login", (&Handlers.AuthHandler{Database: db}).Login).Methods("POST")
	mux.HandleFunc("/signUp", (&Handlers.AuthHandler{Database: db}).SignUp).Methods("POST")
	mux.HandleFunc("/client", (&Handlers.ClientHandler{Database: db}).GetClient).Methods("GET")
	mux.HandleFunc("/client/userid", (&Handlers.ClientHandler{Database: db}).GetClientInfoByUserID).Methods("GET")
	mux.HandleFunc("/clientname", (&Handlers.ClientHandler{Database: db}).GetClientName).Methods("GET")
	mux.HandleFunc("/clientname/userid", (&Handlers.ClientHandler{Database: db}).GetClientNameByUserID).Methods("GET")
	mux.HandleFunc("/clientnames", (&Handlers.ClientHandler{Database: db}).GetAllClientNames).Methods("GET")
	mux.HandleFunc("/clinician/userid", (&Handlers.ClinicianHandler{Database: db}).GetClinicianInfoByUserID).Methods("GET")
	mux.HandleFunc("/clinicianname", (&Handlers.ClinicianHandler{Database: db}).GetClinicianName).Methods("GET")
	mux.HandleFunc("/clinicianname/userid", (&Handlers.ClinicianHandler{Database: db}).GetClinicianNameByUserID).Methods("GET")
	mux.HandleFunc("/cliniciannames", (&Handlers.ClinicianHandler{Database: db}).GetClinicianNames).Methods("GET")
	mux.HandleFunc("/safetyplan", dbHandler.getSafetyPlan).Methods("GET")
	mux.HandleFunc("/client/safetyplan", (&Handlers.SafetyPlanHandler{Database: db}).ClientGetSafetyPlan).Methods("GET")
	mux.HandleFunc("/client/appointments", (&Handlers.AppointmentHandler{Database: db}).ClientGetAppointments).Methods("GET")
	mux.HandleFunc("/clinician/clients", dbHandler.getClients).Methods("GET")
	mux.HandleFunc("/clinician/clinicians", (&Handlers.ClinicianHandler{Database: db}).GetClinicians).Methods("GET")
	mux.HandleFunc("/clinician/safetyplan", (&Handlers.SafetyPlanHandler{Database: db}).ClinicianGetSafetyPlan).Methods("GET")
	mux.HandleFunc("/clinician/safetyplans", (&Handlers.SafetyPlanHandler{Database: db}).ClinicianGetSafetyPlans).Methods("GET")
	mux.HandleFunc("/clinician/safetyplan/add", (&Handlers.SafetyPlanHandler{Database: db}).ClinicianAddSafetyPlan).Methods("POST")
	mux.HandleFunc("/clinician/safetyplan/update", (&Handlers.SafetyPlanHandler{Database: db}).ClinicianUpdateSafetyPlan).Methods("POST")
	mux.HandleFunc("/clinician/safetyplan/delete", (&Handlers.SafetyPlanHandler{Database: db}).ClinicianDeleteSafetyPlan).Methods("POST")
	mux.HandleFunc("/appointment", (&Handlers.AppointmentHandler{Database: db}).GetAppointment).Methods("GET")
	mux.HandleFunc("/appointment/add", (&Handlers.AppointmentHandler{Database: db}).AddAppointment).Methods("POST")
	mux.HandleFunc("/appointment/update", (&Handlers.AppointmentHandler{Database: db}).UpdateAppointment).Methods("POST")
	mux.HandleFunc("/appointment/delete", (&Handlers.AppointmentHandler{Database: db}).DeleteAppointment).Methods("POST")
	mux.HandleFunc("/clinician/appointments", (&Handlers.AppointmentHandler{Database: db}).ClinicianGetAllAppointments).Methods("GET")

	// Allow CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	// Server Configurations
	router := c.Handler(mux)
	log.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", handlers.CORS(originsOK, headersOK, methodsOK)(router))
	log.Fatal(err)
}

// This function gets clients only if you are a clinician
func (db *Database) getClients(w http.ResponseWriter, r *http.Request) {
	claims, er := Auth.IsAuthorized(w, r)
	if er == false {
		return
	}
	var role string = fmt.Sprintf("%v", claims["role"])
	// Check if the person is a clinician
	if role == "2" {
		clinician := Facade.NewClinicianFacade(db.database)
		clients, _ := clinician.GetAllClients()
		b, erro := json.Marshal(clients)
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

func (db *Database) getSafetyPlan(w http.ResponseWriter, r *http.Request) {
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
	intRole, err := strconv.Atoi(role)
	if (role == "1" && userID == fmt.Sprintf("%v", claims["userID"])) || role == "2" {
		person := Facade.NewPersonFacade(db.database)
		safetyPlan, _ := person.GetSafetyPlansByUserID(intUserID, intRole)
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
