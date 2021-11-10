package model

import "fmt"

type Person struct {
	UserID int `db:"userID"`
	//UserID		int 	`db"UserID"`
	UserName    string `db:"username"`
	Password    string `db:"password"`
	FirstName   string `db:"firstname"`
	LastName    string `db:"lastname"`
	Email       string `db:"email"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phonenumber"`
	Role        string `db:"role"`
	Expiration  string `db:"expiration"`
}

func NewPerson(userName string, password string, firstName string, lastName string, email string, address string, phoneNumber string, role string, expiration string) *Person {
	return &Person{UserName: userName, Password: password, FirstName: firstName, LastName: lastName, Email: email, Address: address, PhoneNumber: phoneNumber, Role: role, Expiration: expiration}
}

func (p *Person) GetUserID() int {
	return p.UserID
}

func (p *Person) SetUserID(userID int) {
	p.UserID = userID
}

func (p *Person) GetUserName() string {
	return p.UserName
}

func (p *Person) SetUserName(userName string) {
	p.UserName = userName
}

func (p *Person) GetPassword() string {
	return p.Password
}

func (p *Person) SetPassword(password string) {
	p.Password = password
}

func (p *Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) SetFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *Person) GetLastName() string {
	return p.LastName
}

func (p *Person) SetLastName(lastName string) {
	p.LastName = lastName
}

func (p *Person) GetEmail() string {
	return p.Email
}

func (p *Person) SetEmail(email string) {
	p.Email = email
}

func (p *Person) GetAddress() string {
	return p.Address
}

func (p *Person) SetAddress(address string) {
	p.Address = address
}

func (p *Person) GetPhoneNumber() string {
	return p.PhoneNumber
}

func (p *Person) SetPhoneNumber(phoneNumber string) {
	p.PhoneNumber = phoneNumber
}

func (p *Person) GetRole() string {
	return p.Role
}

func (p *Person) SetRole(role string) {
	p.Role = role
}

func (p *Person) GetExpiration() string {
	return p.Expiration
}

func (p *Person) SetExpiration(expiration string) {
	p.Expiration = expiration
}

func (p Person) Error() string {
	panic("implement me")
}

func (p *Person) Print() string {
	var pString = fmt.Sprintf("UserID: %d\nUsername: %s\nPassword: %s\nFirst Name: %s\nLast Name: %s\nEmail: %s\nAddress: %s\nPhone Number: %s\nRole: %s\nExpiration: %s\n", p.GetUserID(), p.GetUserName(), p.GetPassword(), p.GetFirstName(), p.GetLastName(), p.GetEmail(), p.GetAddress(), p.GetPhoneNumber(), p.GetRole(), p.GetExpiration())
	return pString
}
