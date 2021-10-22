package model

type FamilyMember struct {
	familyID     int
	relationship string
	clientID     int
	userID       int
}

func NewFamilyMember(relationship string, clientID int, userID int) *FamilyMember {
	return &FamilyMember{relationship: relationship, clientID: clientID, userID: userID}
}

func (f *FamilyMember) FamilyID() int {
	return f.familyID
}

func (f *FamilyMember) SetFamilyID(familyID int) {
	f.familyID = familyID
}

func (f *FamilyMember) Relationship() string {
	return f.relationship
}

func (f *FamilyMember) SetRelationship(relationship string) {
	f.relationship = relationship
}

func (f *FamilyMember) ClientID() int {
	return f.clientID
}

func (f *FamilyMember) SetClientID(clientID int) {
	f.clientID = clientID
}

func (f *FamilyMember) UserID() int {
	return f.userID
}

func (f *FamilyMember) SetUserID(userID int) {
	f.userID = userID
}
