package domain

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================================
// PatientAddress
// ============================================================================

type PatientAddressProps struct {
	PatientID string
	PlaceID   string
	Line1     string
	City      string
	State     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PatientAddress struct {
	patientID string
	placeID   string
	line1     string
	city      string
	state     string
	createdAt time.Time
	updatedAt time.Time
}

func NewPatientAddress(
	patientID string,
	line1 string,
	city string,
	state string,
	createdAt time.Time,
) *PatientAddress {
	return &PatientAddress{
		patientID: patientID,
		placeID:   uuid.NewString(),
		line1:     line1,
		city:      city,
		state:     state,
		createdAt: createdAt,
		updatedAt: createdAt,
	}
}

func ReconstitutePatientAddress(p PatientAddressProps) *PatientAddress {
	return &PatientAddress{
		patientID: p.PatientID,
		placeID:   p.PlaceID,
		line1:     p.Line1,
		city:      p.City,
		state:     p.State,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (pa *PatientAddress) PatientID() string    { return pa.patientID }
func (pa *PatientAddress) PlaceID() string      { return pa.placeID }
func (pa *PatientAddress) Line1() string        { return pa.line1 }
func (pa *PatientAddress) City() string         { return pa.city }
func (pa *PatientAddress) State() string        { return pa.state }
func (pa *PatientAddress) CreatedAt() time.Time { return pa.createdAt }
func (pa *PatientAddress) UpdatedAt() time.Time { return pa.updatedAt }

func (pa *PatientAddress) SetLine1(line1 string) *PatientAddress {
	pa.line1 = line1
	return pa
}

func (pa *PatientAddress) SetCity(city string) *PatientAddress {
	pa.city = city
	return pa
}

func (pa *PatientAddress) SetState(state string) *PatientAddress {
	pa.state = state
	return pa
}

func (pa *PatientAddress) SetUpdatedAt(updatedAt time.Time) *PatientAddress {
	pa.updatedAt = updatedAt
	return pa
}
