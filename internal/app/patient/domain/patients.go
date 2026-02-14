package domain

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================================
// Patient
// ============================================================================

type PatientProps struct {
	PatientID string
	FirstName string
	LastName  string
	Gender    string
	Dob       time.Time
	StaffID   *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Patient struct {
	patientID string
	firstName string
	lastName  string
	gender    string
	dob       time.Time
	staffID   *string
	createdAt time.Time
	updatedAt time.Time
}

func NewPatient(
	firstName string,
	lastName string,
	gender string,
	dob time.Time,
	staffID *string,
	createdAt time.Time,
) *Patient {
	return &Patient{
		patientID: uuid.NewString(),
		firstName: firstName,
		lastName:  lastName,
		gender:    gender,
		dob:       dob,
		staffID:   staffID,
		createdAt: createdAt,
		updatedAt: createdAt,
	}
}

func ReconstitutePatient(p PatientProps) *Patient {
	return &Patient{
		patientID: p.PatientID,
		firstName: p.FirstName,
		lastName:  p.LastName,
		gender:    p.Gender,
		dob:       p.Dob,
		staffID:   p.StaffID,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (p *Patient) PatientID() string     { return p.patientID }
func (p *Patient) FirstName() string     { return p.firstName }
func (p *Patient) LastName() string      { return p.lastName }
func (p *Patient) Gender() string        { return p.gender }
func (p *Patient) Dob() time.Time        { return p.dob }
func (p *Patient) StaffID() *string      { return p.staffID }
func (p *Patient) CreatedAt() time.Time  { return p.createdAt }
func (p *Patient) UpdatedAt() time.Time  { return p.updatedAt }

func (p *Patient) SetFirstName(firstName string) *Patient {
	p.firstName = firstName
	return p
}

func (p *Patient) SetLastName(lastName string) *Patient {
	p.lastName = lastName
	return p
}

func (p *Patient) SetGender(gender string) *Patient {
	p.gender = gender
	return p
}

func (p *Patient) SetDob(dob time.Time) *Patient {
	p.dob = dob
	return p
}

func (p *Patient) SetStaffID(staffID *string) *Patient {
	p.staffID = staffID
	return p
}

func (p *Patient) SetUpdatedAt(updatedAt time.Time) *Patient {
	p.updatedAt = updatedAt
	return p
}
