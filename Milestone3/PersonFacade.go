package facade

type PersonFacade struct {
	userID int
	userName string
	password string
	name string
	address string
	phoneNumber string
}

func newPersonFacade(userID int, userName string, password string, name string, address string, phoneNumber string) *PersonFacade {
	return &PersonFacade{
		userID,
		userName,
		password,
		name,
		address,
		phoneNumber}
}

func (p *PersonFacade) updatePassword(userID int, password string) error {
	p.password = password
	return nil
}
