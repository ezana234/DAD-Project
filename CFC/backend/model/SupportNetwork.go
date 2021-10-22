package model

type SupportNetwork struct {
	supportID    int
	relationship string
	phoneNumber  string
	clientID     int
}

func NewSupportNetwork(relationship string, phoneNumber string, clientID int) *SupportNetwork {
	return &SupportNetwork{relationship: relationship, phoneNumber: phoneNumber, clientID: clientID}
}

func (s *SupportNetwork) SupportID() int {
	return s.supportID
}

func (s *SupportNetwork) SetSupportID(supportID int) {
	s.supportID = supportID
}

func (s *SupportNetwork) Relationship() string {
	return s.relationship
}

func (s *SupportNetwork) SetRelationship(relationship string) {
	s.relationship = relationship
}

func (s *SupportNetwork) PhoneNumber() string {
	return s.phoneNumber
}

func (s *SupportNetwork) SetPhoneNumber(phoneNumber string) {
	s.phoneNumber = phoneNumber
}

func (s *SupportNetwork) ClientID() int {
	return s.clientID
}

func (s *SupportNetwork) SetClientID(clientID int) {
	s.clientID = clientID
}
