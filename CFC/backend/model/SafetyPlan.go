package model

import "fmt"

type SafetyPlan struct {
	SafetyID             int
	Triggers             string
	WarningSigns         string
	DestructiveBehaviors string
	InternalStrategies   string
	UpdatedClinician     int
	UpdatedDatetime      string
	ClientID             int
	ClinicianID          int
}

func NewSafetyPlan(triggers string, warningSigns string, destructiveBehaviors string, internalStrategies string, updatedDatetime string, updatedClinician int, clientID int, clinicianID int) *SafetyPlan {
	return &SafetyPlan{Triggers: triggers, WarningSigns: warningSigns, DestructiveBehaviors: destructiveBehaviors, InternalStrategies: internalStrategies, UpdatedDatetime: updatedDatetime, UpdatedClinician: updatedClinician, ClientID: clientID, ClinicianID: clinicianID}
}

func (sp *SafetyPlan) GetSafetyID() int {
	return sp.SafetyID
}

func (sp *SafetyPlan) SetSafetyID(safetyID int) {
	sp.SafetyID = safetyID
}

func (sp *SafetyPlan) GetTriggers() string {
	return sp.Triggers
}

func (sp *SafetyPlan) SetTriggers(triggers string) {
	sp.Triggers = triggers
}

func (sp *SafetyPlan) GetWarningSigns() string {
	return sp.WarningSigns
}

func (sp *SafetyPlan) SetWarningSigns(warningSigns string) {
	sp.WarningSigns = warningSigns
}

func (sp *SafetyPlan) GetDestructiveBehaviors() string {
	return sp.DestructiveBehaviors
}

func (sp *SafetyPlan) SetDestructiveBehaviors(destructiveBehaviors string) {
	sp.DestructiveBehaviors = destructiveBehaviors
}

func (sp *SafetyPlan) GetInternalStrategies() string {
	return sp.InternalStrategies
}

func (sp *SafetyPlan) SetInternalStrategies(internalStrategies string) {
	sp.InternalStrategies = internalStrategies
}

func (sp *SafetyPlan) GetUpdatedDatetime() string {
	return sp.UpdatedDatetime
}

func (sp *SafetyPlan) SetUpdatedDatetime(updatedDatetime string) {
	sp.UpdatedDatetime = updatedDatetime
}

func (sp *SafetyPlan) GetUpdatedClinician() int {
	return sp.UpdatedClinician
}

func (sp *SafetyPlan) SetUpdatedClinician(updatedClinician int) {
	sp.UpdatedClinician = updatedClinician
}

func (sp *SafetyPlan) GetClientID() int {
	return sp.ClientID
}

func (sp *SafetyPlan) SetClientID(clientID int) {
	sp.ClientID = clientID
}

func (sp *SafetyPlan) GetClinicianID() int {
	return sp.ClinicianID
}

func (sp *SafetyPlan) SetClinicianID(clinicianID int) {
	sp.ClinicianID = clinicianID
}

func (sp *SafetyPlan) SafetyPlanToString() string {
	spStr := fmt.Sprintf("SafetyID: %d, Triggers: %s, WarningSigns: %s, DestructiveBehaviors: %s, InternalStrategies: %s, UpdatedDateTime: %s, UpdatedClinician: %d, ClientID: %d, ClinicianID: %d", sp.SafetyID, sp.Triggers, sp.WarningSigns, sp.InternalStrategies, sp.UpdatedDatetime, sp.UpdatedClinician, sp.ClientID, sp.ClinicianID)
	return spStr
}
