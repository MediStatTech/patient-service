package repo

import (
	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	"github.com/google/uuid"
)

// ============================================================================
// Patients Mappers
// ============================================================================

func toPatientProps(patient Patient) domain.PatientProps {
	var staffID *string
	if patient.StaffID.Valid {
		s := patient.StaffID.UUID.String()
		staffID = &s
	}

	return domain.PatientProps{
		PatientID: patient.PatientID.String(),
		FirstName: patient.FirstName,
		LastName:  patient.LastName,
		Gender:    patient.Gender,
		Dob:       patient.Dob,
		StaffID:   staffID,
		CreatedAt: patient.CreatedAt,
		UpdatedAt: patient.UpdatedAt,
	}
}

func patientToCreateParams(patient *domain.Patient) []any {
	id, _ := uuid.Parse(patient.PatientID())
	
	var staffID uuid.NullUUID
	if patient.StaffID() != nil {
		parsedStaffID, _ := uuid.Parse(*patient.StaffID())
		staffID = uuid.NullUUID{UUID: parsedStaffID, Valid: true}
	}

	return []any{
		id,
		patient.FirstName(),
		patient.LastName(),
		patient.Gender(),
		patient.Dob(),
		staffID,
		patient.CreatedAt(),
		patient.UpdatedAt(),
	}
}

func patientToUpdateParams(patient *domain.Patient) []any {
	id, _ := uuid.Parse(patient.PatientID())
	
	var staffID uuid.NullUUID
	if patient.StaffID() != nil {
		parsedStaffID, _ := uuid.Parse(*patient.StaffID())
		staffID = uuid.NullUUID{UUID: parsedStaffID, Valid: true}
	}

	return []any{
		id,
		patient.FirstName(),
		patient.LastName(),
		patient.Gender(),
		patient.Dob(),
		staffID,
		patient.UpdatedAt(),
	}
}

// ============================================================================
// Patient Addresses Mappers
// ============================================================================

func toPatientAddressProps(address PatientAddress) domain.PatientAddressProps {
	return domain.PatientAddressProps{
		PatientID: address.PatientID.String(),
		PlaceID:   address.PlaceID.String(),
		Line1:     address.Line1,
		City:      address.City,
		State:     address.State,
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
	}
}

func patientAddressToCreateParams(address *domain.PatientAddress) []any {
	patientID, _ := uuid.Parse(address.PatientID())
	placeID, _ := uuid.Parse(address.PlaceID())

	return []any{
		patientID,
		placeID,
		address.Line1(),
		address.City(),
		address.State(),
		address.CreatedAt(),
		address.UpdatedAt(),
	}
}

func patientAddressToUpdateParams(address *domain.PatientAddress) []any {
	placeID, _ := uuid.Parse(address.PlaceID())

	return []any{
		placeID,
		address.Line1(),
		address.City(),
		address.State(),
		address.UpdatedAt(),
	}
}

// ============================================================================
// Patient Contact Infos Mappers
// ============================================================================

func toPatientContactInfoProps(contactInfo PatientContactInfo) *domain.PatientContactInfoProps {
	return &domain.PatientContactInfoProps{
		PatientID: contactInfo.PatientID.String(),
		ContactID: contactInfo.ContactID.String(),
		Phone:     contactInfo.Phone,
		Email:     contactInfo.Email,
		Primary:   contactInfo.Primary,
		CreatedAt: contactInfo.CreatedAt,
		UpdatedAt: contactInfo.UpdatedAt,
	}
}

func patientContactInfoToCreateParams(contactInfo *domain.PatientContactInfo) []any {
	patientID, _ := uuid.Parse(contactInfo.PatientID())
	contactID, _ := uuid.Parse(contactInfo.ContactID())

	return []any{
		patientID,
		contactID,
		contactInfo.Phone(),
		contactInfo.Email(),
		contactInfo.Primary(),
		contactInfo.CreatedAt(),
		contactInfo.UpdatedAt(),
	}
}

func patientContactInfoToUpdateParams(contactInfo *domain.PatientContactInfo) []any {
	contactID, _ := uuid.Parse(contactInfo.ContactID())

	return []any{
		contactID,
		contactInfo.Phone(),
		contactInfo.Email(),
		contactInfo.Primary(),
		contactInfo.UpdatedAt(),
	}
}

// ============================================================================
// Patient Diseases Mappers
// ============================================================================

func toPatientDiseasProps(patientDiseas PatientDiseas) domain.PatientDiseasProps {
	return domain.PatientDiseasProps{
		PatientID: patientDiseas.PatientID.String(),
		DiseasID:  patientDiseas.DiseasID.String(),
		CreatedAt: patientDiseas.CreatedAt,
		UpdatedAt: patientDiseas.UpdatedAt,
	}
}

func patientDiseasToCreateParams(patientDiseas *domain.PatientDiseas) []any {
	patientID, _ := uuid.Parse(patientDiseas.PatientID())
	diseasID, _ := uuid.Parse(patientDiseas.DiseasID())

	return []any{
		patientID,
		diseasID,
		patientDiseas.CreatedAt(),
		patientDiseas.UpdatedAt(),
	}
}

func patientDiseasToUpdateParams(patientDiseas *domain.PatientDiseas) []any {
	patientID, _ := uuid.Parse(patientDiseas.PatientID())
	diseasID, _ := uuid.Parse(patientDiseas.DiseasID())

	return []any{
		patientID,
		diseasID,
		patientDiseas.UpdatedAt(),
	}
}
