package domain

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================================
// PatientContactInfo
// ============================================================================

type PatientContactInfoProps struct {
	PatientID string
	ContactID string
	Phone     string
	Email     string
	Primary   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PatientContactInfo struct {
	patientID string
	contactID string
	phone     string
	email     string
	primary   bool
	createdAt time.Time
	updatedAt time.Time
}

func NewPatientContactInfo(
	patientID string,
	phone string,
	email string,
	primary bool,
	createdAt time.Time,
) *PatientContactInfo {
	return &PatientContactInfo{
		patientID: patientID,
		contactID: uuid.NewString(),
		phone:     phone,
		email:     email,
		primary:   primary,
		createdAt: createdAt,
		updatedAt: createdAt,
	}
}

func ReconstitutePatientContactInfo(p PatientContactInfoProps) *PatientContactInfo {
	return &PatientContactInfo{
		patientID: p.PatientID,
		contactID: p.ContactID,
		phone:     p.Phone,
		email:     p.Email,
		primary:   p.Primary,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (pci *PatientContactInfo) PatientID() string    { return pci.patientID }
func (pci *PatientContactInfo) ContactID() string    { return pci.contactID }
func (pci *PatientContactInfo) Phone() string        { return pci.phone }
func (pci *PatientContactInfo) Email() string        { return pci.email }
func (pci *PatientContactInfo) Primary() bool        { return pci.primary }
func (pci *PatientContactInfo) CreatedAt() time.Time { return pci.createdAt }
func (pci *PatientContactInfo) UpdatedAt() time.Time { return pci.updatedAt }

func (pci *PatientContactInfo) SetPhone(phone string) *PatientContactInfo {
	pci.phone = phone
	return pci
}

func (pci *PatientContactInfo) SetEmail(email string) *PatientContactInfo {
	pci.email = email
	return pci
}

func (pci *PatientContactInfo) SetPrimary(primary bool) *PatientContactInfo {
	pci.primary = primary
	return pci
}

func (pci *PatientContactInfo) SetUpdatedAt(updatedAt time.Time) *PatientContactInfo {
	pci.updatedAt = updatedAt
	return pci
}
