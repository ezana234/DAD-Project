package main

import (
	"CFC/backend/CFC/backend/DB"
	// "strings"
	// "time"

	// 	Auth "CFC/backend/CFC/backend/auth"
	// 	Facade "CFC/backend/CFC/backend/facade"
	// 	Handlers "CFC/backend/CFC/backend/handlers"

	Auth "CFC/backend/CFC/backend/auth"
	"CFC/backend/CFC/backend/facade"
	Facade "CFC/backend/CFC/backend/facade"
	Handlers "CFC/backend/CFC/backend/handlers"
	"CFC/backend/CFC/backend/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	// Routes
	// mux.HandleFunc("/login", dbHandler.login).Methods("POST")
	mux.HandleFunc("/login", (&Handlers.AuthHandler{Database: db}).Login).Methods("POST")
	// 	mux.HandleFunc("/client", dbHandler.client).Methods("GET")
	// 	mux.HandleFunc("/person", (&Handlers.PersonHandler{Database: db}).GetPerson).Methods("GET")
	// mux.HandleFunc("/login", dbHandler.login).Methods("POST")
	mux.HandleFunc("/signUp", dbHandler.signUp).Methods("POST")
	mux.HandleFunc("/client", dbHandler.getClient).Methods("GET")
	mux.HandleFunc("/clinician/clients", dbHandler.getClients).Methods("GET")
	// mux.HandleFunc("/clinician/safetyplan")
	mux.HandleFunc("/client/safetyplan", (&Handlers.SafetyPlanHandler{Database: db}).ClientGetSafetyPlan).Methods("GET")

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

// func (db *Database) login(w http.ResponseWriter, r *http.Request) {
// 	type Login struct {
// 		Email    string
// 		Password string
// 	}
// 	var logStruct Login
// 	body := json.NewDecoder(r.Body).Decode(&logStruct)
// 	if body != nil {
// 		http.Error(w, body.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	person := Facade.NewPersonFacade(db.database)
// 	pers := person.GetPersonByEmail(logStruct.Email, logStruct.Password)
// 	if pers.GetUserID() == 0 {
// 		http.Error(w, "Bad Login", http.StatusUnauthorized)
// 		return
// 	} else {
// 		tokenString, err := Auth.GenerateJWT(pers.GetUserID(), pers.GetEmail(), pers.GetRole())
// 		println(tokenString)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		resp := make(map[string]string)
// 		resp["token"] = tokenString
// 		b, err := json.Marshal(resp)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(b)
// 	}
// }

// func (db *Database) client(w http.ResponseWriter, r *http.Request) {
// 	claims, er := Auth.IsAuthorized(w, r)

// This returns a jwt upon a successful login
func (db *Database) signUp(w http.ResponseWriter, r *http.Request) {
	type sign struct {
		Username    string
		FirstName   string
		LastName    string
		Email       string
		Address     string
		Password    string
		PhoneNumber string
		DOB         string
	}
	var signStruct sign
	body := json.NewDecoder(r.Body).Decode(&signStruct)
	if body != nil {
		http.Error(w, body.Error(), http.StatusBadRequest)
		return
	}

	person := Facade.NewPersonFacade(db.database)
	newPers := model.NewPerson(
		signStruct.Username,
		signStruct.Password,
		signStruct.FirstName,
		signStruct.LastName,
		signStruct.Email,
		signStruct.Address,
		signStruct.PhoneNumber,
		"1",
		" ",
		signStruct.DOB)
	userID, err := person.CreateNewPerson(*newPers)
	if err != 1 {
		http.Error(w, "Couldn't Create Person", http.StatusBadRequest)
		return
	} else {
		// fmt.Println(userID)
		client := facade.NewClientFacade(db.database)
		clientModel := model.NewClient(userID)
		client.AddClient(*clientModel)
		resp := make(map[string]string)
		resp["message"] = "Client added to Database"
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
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
		tokenString, erro := Auth.GenerateJWT(pers.GetUserID(), pers.GetEmail(), pers.GetRole())
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	claims, er := Auth.IsAuthorized(w, r)
	if er == false {
		return
	}
	person := Facade.NewPersonFacade(db.database)
	var userID int = int(claims["userID"].(float64))
	pers, err := person.GetPerson(userID)
	if err == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
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

//func GenerateJWT(userID int, email string, role string) (string, error) {
//	var mySigningKey = []byte("CFC-Secret8")
//	token := jwt.New(jwt.SigningMethodHS256)
//	claims := token.Claims.(jwt.MapClaims)
//
//	claims["authorized"] = true
//	claims["userID"] = userID
//	claims["email"] = email
//	claims["role"] = role
//	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
//
//	tokenString, err := token.SignedString(mySigningKey)
//
//	if err != nil {
//		fmt.Errorf("Something Went Wrong: %s", err.Error())
//		return "", err
//	}
//
//	return tokenString, nil
//}

//func isAuthorized(w http.ResponseWriter, r *http.Request) (jwt.MapClaims, bool) {
//	fmt.Println(r.Header)
//	if r.Header["Authorization"] == nil {
//		resp := make(map[string]string)
//		resp["error"] = "No Token Found"
//		json.NewEncoder(w).Encode(resp)
//		return nil, false
//	}
//
//	var mySigningKey = []byte("CFC-Secret8")
//
//	token, err := jwt.Parse(strings.Split(r.Header["Authorization"][0], " ")[1], func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("There was an error in parsing")
//		}
//		return mySigningKey, nil
//	})
//
//	if err != nil {
//		resp := make(map[string]string)
//		resp["error"] = "Your Token is invalid."
//		json.NewEncoder(w).Encode(resp)
//		return nil, false
//	}
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		fmt.Println(claims)
//		// if claims["role"] == "admin" {
//
//		// 	r.Header.Set("Role", "admin")
//		// 	handler.ServeHTTP(w, r)
//		// 	return
//
//		// } else if claims["role"] == "user" {
//
//		// 	r.Header.Set("Role", "user")
//		// 	handler.ServeHTTP(w, r)
//		// 	return
//		// }
//		return claims, true
//	} else {
//		return nil, false
//	}
//
//}

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

// Generates a JWT
// func GenerateJWT(userID int, email string, role string) (string, error) {
// 	var mySigningKey = []byte("CFC-Secret8")
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["userID"] = userID
// 	claims["email"] = email
// 	claims["role"] = role
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

// 	tokenString, err := token.SignedString(mySigningKey)

// 	if err != nil {
// 		fmt.Errorf("Something Went Wrong: %s", err.Error())
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// Checks if the JWT is valid
// func isAuthorized(w http.ResponseWriter, r *http.Request) (jwt.MapClaims, bool) {
// 	fmt.Println(r.Header)
// 	if r.Header["Authorization"] == nil {
// 		resp := make(map[string]string)
// 		resp["error"] = "No Token Found"
// 		json.NewEncoder(w).Encode(resp)
// 		return nil, false
// 	}

// 	var mySigningKey = []byte("CFC-Secret8")

// 	token, err := jwt.Parse(strings.Split(r.Header["Authorization"][0], " ")[1], func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("There was an error in parsing")
// 		}
// 		return mySigningKey, nil
// 	})

// 	if err != nil {
// 		resp := make(map[string]string)
// 		resp["error"] = "Your Token is invalid."
// 		json.NewEncoder(w).Encode(resp)
// 		return nil, false
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims, true
// 	} else {
// 		return nil, false
// 	}

// }
