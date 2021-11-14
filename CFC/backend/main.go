package main

import (
	"CFC/backend/CFC/backend/DB"
	Facade "CFC/backend/CFC/backend/facade"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

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
	// Routes
	mux.HandleFunc("/login", dbHandler.login).Methods("POST")
	mux.HandleFunc("/client", dbHandler.getClient).Methods("GET")
	mux.HandleFunc("/clinician/clients", dbHandler.getClients).Methods("GET")
	// Allow CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://34.227.30.182:3000"}, //you service is available and allowed for this base url
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

	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	// Server Configurations
	router := c.Handler(mux)
	log.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", handlers.CORS(originsOK, headersOK, methodsOK)(router))
	log.Fatal(err)
}

// This returns a jwt upon a successful login
func (db *Database) login(w http.ResponseWriter, r *http.Request) {
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
	if pers.GetUserID() == 0 {
		http.Error(w, "Bad Login", http.StatusUnauthorized)
		return
	} else {
		tokenString, erro := GenerateJWT(pers.GetUserID(), pers.GetEmail(), pers.GetRole())
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
		resp := make(map[string]string)
		resp["token"] = tokenString
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

// This function gets the client from the JWT
func (db *Database) getClient(w http.ResponseWriter, r *http.Request) {
	claims, er := isAuthorized(w, r)
	if er == false {
		return
	}
	person := Facade.NewPersonFacade(db.database)
	var userID int = int(claims["userID"].(float64))
	pers, err := person.GetPerson(userID)
	if err == 0 {
		http.Error(w, pers.Error(), http.StatusNotFound)
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
}

// This function gets clients only if you are a clinician
func (db *Database) getClients(w http.ResponseWriter, r *http.Request) {
	claims, er := isAuthorized(w, r)
	if er == false {
		return
	}
	var role string = fmt.Sprintf("%v", claims["role"])
	// Check if the person is a clinician
	if role == "2" {
		clinician := Facade.NewClinicianFacade(db.database)
		clients := clinician.GetAllClients()
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

// Generates a JWT
func GenerateJWT(userID int, email string, role string) (string, error) {
	var mySigningKey = []byte("CFC-Secret8")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userID"] = userID
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

// Checks if the JWT is valid
func isAuthorized(w http.ResponseWriter, r *http.Request) (jwt.MapClaims, bool) {
	if r.Header["Authorization"] == nil {
		resp := make(map[string]string)
		resp["error"] = "No Token Found"
		json.NewEncoder(w).Encode(resp)
		return nil, false
	}

	var mySigningKey = []byte("CFC-Secret8")

	token, err := jwt.Parse(strings.Split(r.Header["Authorization"][0], " ")[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		resp := make(map[string]string)
		resp["error"] = "Your Token is invalid."
		json.NewEncoder(w).Encode(resp)
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}

}
