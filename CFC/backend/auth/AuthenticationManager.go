package auth

import (
	Model "CFC/backend/CFC/backend/model"
)

type AuthenticationManager struct {
	user *Model.Person
}

func NewAuthenticationManager() *AuthenticationManager {
	return &AuthenticationManager{}
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

func (am *AuthenticationManager) IsCurrentUserClinician() bool {
	return am.user.GetRole() == "2"
}

func (am *AuthenticationManager) IsCurrentUserFamilyMember() bool {
	return am.user.GetRole() == "3"
}

func (am *AuthenticationManager) IsCurrentUserAdmin() bool {
	return am.user.GetRole() == "4"
}

func (am *AuthenticationManager) IsCurrentUser(userID int) bool {
	return am.user.GetUserID() == userID
}
