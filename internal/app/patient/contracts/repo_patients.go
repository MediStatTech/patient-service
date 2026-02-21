package contracts

import (
	"context"

	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

// ============================================================================
// Patients Repository
// ============================================================================

type PatientsRepo interface {
	FindAll(ctx context.Context) ([]domain.PatientProps, error)
	FindByID(ctx context.Context, patientID string) (domain.PatientProps, error)
	FindByStaffID(ctx context.Context, staffID string) ([]domain.PatientProps, error)
	CreateMut(patient *domain.Patient) *postgres.Mutation
	UpdateMut(patient *domain.Patient) *postgres.Mutation
	CreateBatchMut(patients []*domain.Patient) []*postgres.Mutation
}

// ============================================================================
// Patient Addresses Repository
// ============================================================================

type PatientAddressesRepo interface {
	FindAll(ctx context.Context) ([]domain.PatientAddressProps, error)
	FindByID(ctx context.Context, placeID string) (domain.PatientAddressProps, error)
	FindByPatientID(ctx context.Context, patientID string) ([]domain.PatientAddressProps, error)
	CreateMut(address *domain.PatientAddress) *postgres.Mutation
	UpdateMut(address *domain.PatientAddress) *postgres.Mutation
	DeleteMut(placeID string) *postgres.Mutation
	CreateBatchMut(addresses []*domain.PatientAddress) []*postgres.Mutation
}

// ============================================================================
// Patient Contact Infos Repository
// ============================================================================

type PatientContactInfosRepo interface {
	FindAll(ctx context.Context) ([]*domain.PatientContactInfoProps, error)
	FindByID(ctx context.Context, contactID string) (*domain.PatientContactInfoProps, error)
	FindByPatientID(ctx context.Context, patientID string) ([]*domain.PatientContactInfoProps, error)
	FindByPatientIDAndPrimary(ctx context.Context, patientID string) (*domain.PatientContactInfoProps, error)
	CreateMut(contactInfo *domain.PatientContactInfo) *postgres.Mutation
	UpdateMut(contactInfo *domain.PatientContactInfo) *postgres.Mutation
	DeleteMut(contactID string) *postgres.Mutation
	CreateBatchMut(contactInfos []*domain.PatientContactInfo) []*postgres.Mutation
}

// ============================================================================
// Patient Diseases Repository
// ============================================================================

type PatientDiseasesRepo interface {
	FindAll(ctx context.Context) ([]domain.PatientDiseasProps, error)
	FindByPatientID(ctx context.Context, patientID string) ([]domain.PatientDiseasProps, error)
	FindByDiseasID(ctx context.Context, diseasID string) ([]domain.PatientDiseasProps, error)
	FindByPatientAndDiseas(ctx context.Context, patientID, diseasID string) (domain.PatientDiseasProps, error)
	CreateMut(patientDiseas *domain.PatientDiseas) *postgres.Mutation
	UpdateMut(patientDiseas *domain.PatientDiseas) *postgres.Mutation
	DeleteMut(patientID, diseasID string) *postgres.Mutation
	CreateBatchMut(patientDiseases []*domain.PatientDiseas) []*postgres.Mutation
}
