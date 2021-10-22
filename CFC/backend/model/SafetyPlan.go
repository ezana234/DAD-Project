package model

type SafetyPlan struct {
	safetyID             int
	triggers             string
	warningSigns         string
	destructiveBehaviors string
	internalStrategies   string
	updatedDatetime      string
	clientID             int
	clinicianID          int
}

func NewSafetyPlan(triggers string, warningSigns string, destructiveBehaviors string, internalStrategies string, updatedDatetime string, clientID int, clinicianID int) *SafetyPlan {
	return &SafetyPlan{triggers: triggers, warningSigns: warningSigns, destructiveBehaviors: destructiveBehaviors, internalStrategies: internalStrategies, updatedDatetime: updatedDatetime, clientID: clientID, clinicianID: clinicianID}
}

func (s *SafetyPlan) SafetyID() int {
	return s.safetyID
}

func (s *SafetyPlan) SetSafetyID(safetyID int) {
	s.safetyID = safetyID
}

func (s *SafetyPlan) Triggers() string {
	return s.triggers
}

func (s *SafetyPlan) SetTriggers(triggers string) {
	s.triggers = triggers
}

func (s *SafetyPlan) WarningSigns() string {
	return s.warningSigns
}

func (s *SafetyPlan) SetWarningSigns(warningSigns string) {
	s.warningSigns = warningSigns
}

func (s *SafetyPlan) DestructiveBehaviors() string {
	return s.destructiveBehaviors
}

func (s *SafetyPlan) SetDestructiveBehaviors(destructiveBehaviors string) {
	s.destructiveBehaviors = destructiveBehaviors
}

func (s *SafetyPlan) InternalStrategies() string {
	return s.internalStrategies
}

func (s *SafetyPlan) SetInternalStrategies(internalStrategies string) {
	s.internalStrategies = internalStrategies
}

func (s *SafetyPlan) UpdatedDatetime() string {
	return s.updatedDatetime
}

func (s *SafetyPlan) SetUpdatedDatetime(updatedDatetime string) {
	s.updatedDatetime = updatedDatetime
}

func (s *SafetyPlan) ClientID() int {
	return s.clientID
}

func (s *SafetyPlan) SetClientID(clientID int) {
	s.clientID = clientID
}

func (s *SafetyPlan) ClinicianID() int {
	return s.clinicianID
}

func (s *SafetyPlan) SetClinicianID(clinicianID int) {
	s.clinicianID = clinicianID
}
