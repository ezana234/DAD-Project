package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

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

func IsAuthorized(w http.ResponseWriter, r *http.Request) (jwt.MapClaims, bool) {
	fmt.Println(r.Header)
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
		fmt.Println(claims)
		// if claims["role"] == "admin" {

		// 	r.Header.Set("Role", "admin")
		// 	handler.ServeHTTP(w, r)
		// 	return

		// } else if claims["role"] == "user" {

		// 	r.Header.Set("Role", "user")
		// 	handler.ServeHTTP(w, r)
		// 	return
		// }
		return claims, true
	} else {
		return nil, false
	}

}
