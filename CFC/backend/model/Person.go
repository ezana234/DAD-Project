package model

type Person struct {
	userID      int
	userName    string
	password    string
	firstName   string
	lastName    string
	email       string
	address     string
	phoneNumber string
	role        string
}

func NewPerson(userName string, password string, firstName string, lastName string, email string, address string, phoneNumber string, role string) *Person {
	return &Person{userName: userName, password: password, firstName: firstName, lastName: lastName, email: email, address: address, phoneNumber: phoneNumber, role: role}
}

func (p *Person) UserID() int {
	return p.userID
}

func (p *Person) SetUserID(userID int) {
	p.userID = userID
}

func (p *Person) UserName() string {
	return p.userName
}

func (p *Person) SetUserName(userName string) {
	p.userName = userName
}

func (p *Person) Password() string {
	return p.password
}

func (p *Person) SetPassword(password string) {
	p.password = password
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}

func (p *Person) LastName() string {
	return p.lastName
}

func (p *Person) SetLastName(lastName string) {
	p.lastName = lastName
}

func (p *Person) Email() string {
	return p.email
}

func (p *Person) SetEmail(email string) {
	p.email = email
}

func (p *Person) Address() string {
	return p.address
}

func (p *Person) SetAddress(address string) {
	p.address = address
}

func (p *Person) PhoneNumber() string {
	return p.phoneNumber
}

func (p *Person) SetPhoneNumber(phoneNumber string) {
	p.phoneNumber = phoneNumber
}

func (p *Person) Role() string {
	return p.role
}

func (p *Person) SetRole(role string) {
	p.role = role
}

func (p Person) Error() string {
	panic("implement me")
}
