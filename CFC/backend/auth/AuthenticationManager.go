package auth

import (
	Model "CFC/backend/CFC/backend/model"
	"encoding/json"
)

type AuthenticationManager struct {
	user        *Model.Person
	clientID    int
	clinicianID int
}

func NewAuthenticationManager() *AuthenticationManager {
	return &AuthenticationManager{}
}

func NewAuthenticationManagerFromJson(inJson []byte) *AuthenticationManager {
	var am = AuthenticationManager{}

	var err = json.Unmarshal(inJson, &am)
	if err != nil {
		panic(err)
	}

	return &am
}

func (am *AuthenticationManager) LoginUser(p *Model.Person) {
	am.user = p
}

func (am *AuthenticationManager) GetCurrentUser() *Model.Person {
	return am.user
}

func (am *AuthenticationManager) IsCurrentUserClient() bool {
	return am.user.GetRole() == "1"
}

func (am *AuthenticationManager) GetCurrentUserClientID() int {
	return am.clientID
}

func (am *AuthenticationManager) SetCurrentUserClientID(clientID int) {
	am.clientID = clientID
}

func (am *AuthenticationManager) IsCurrentUserClinician() bool {
	return am.user.GetRole() == "2"
}

func (am *AuthenticationManager) GetCurrentUserClinicianID() int {
	return am.clinicianID
}

func (am *AuthenticationManager) SetCurrentUserClinicianID(clinicianID int) {
	am.clinicianID = clinicianID
}

func (am *AuthenticationManager) IsCurrentUserAdmin() bool {
	return am.user.GetRole() == "3"
}

func (am *AuthenticationManager) IsCurrentUser(userID int) bool {
	return am.user.GetUserID() == userID
}

func (am *AuthenticationManager) ToJSON() []byte {
	amJson, err := json.Marshal(am)
	if err != nil {
		panic(err)
	}

	return amJson
}
