package domain

import "time"

// ============================================================================
// PatientDiseas
// ============================================================================

type PatientDiseasProps struct {
	PatientID string
	DiseasID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PatientDiseas struct {
	patientID string
	diseasID  string
	createdAt time.Time
	updatedAt time.Time
}

func NewPatientDiseas(
	patientID string,
	diseasID string,
	createdAt time.Time,
) *PatientDiseas {
	return &PatientDiseas{
		patientID: patientID,
		diseasID:  diseasID,
		createdAt: createdAt,
		updatedAt: createdAt,
	}
}

func ReconstitutePatientDiseas(p PatientDiseasProps) *PatientDiseas {
	return &PatientDiseas{
		patientID: p.PatientID,
		diseasID:  p.DiseasID,
		createdAt: p.CreatedAt,
		updatedAt: p.UpdatedAt,
	}
}

func (pd *PatientDiseas) PatientID() string    { return pd.patientID }
func (pd *PatientDiseas) DiseasID() string     { return pd.diseasID }
func (pd *PatientDiseas) CreatedAt() time.Time { return pd.createdAt }
func (pd *PatientDiseas) UpdatedAt() time.Time { return pd.updatedAt }

func (pd *PatientDiseas) SetUpdatedAt(updatedAt time.Time) *PatientDiseas {
	pd.updatedAt = updatedAt
	return pd
}
