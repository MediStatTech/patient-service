package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

// ============================================================================
// Patients Repository
// ============================================================================

type PatientsRepository struct {
	queries *Queries
}

func NewPatientsRepository(db *sql.DB) *PatientsRepository {
	return &PatientsRepository{
		queries: New(db),
	}
}

func (r *PatientsRepository) FindAll(ctx context.Context) ([]domain.PatientProps, error) {
	patients, err := r.queries.ListPatients(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PatientProps, 0, len(patients))
	for _, patient := range patients {
		result = append(result, toPatientProps(patient))
	}

	return result, nil
}

func (r *PatientsRepository) FindByID(ctx context.Context, patientID string) (domain.PatientProps, error) {
	id, err := uuid.Parse(patientID)
	if err != nil {
		return domain.PatientProps{}, err
	}

	patient, err := r.queries.GetPatient(ctx, id)
	if err != nil {
		return domain.PatientProps{}, err
	}

	return toPatientProps(patient), nil
}

func (r *PatientsRepository) CreateMut(patient *domain.Patient) *postgres.Mutation {
	return postgres.NewMutation(
		CreatePatient,
		patientToCreateParams(patient)...,
	)
}

func (r *PatientsRepository) UpdateMut(patient *domain.Patient) *postgres.Mutation {
	return postgres.NewMutation(
		UpdatePatient,
		patientToUpdateParams(patient)...,
	)
}

func (r *PatientsRepository) CreateBatchMut(patients []*domain.Patient) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(patients))
	for _, patient := range patients {
		mutations = append(mutations, r.CreateMut(patient))
	}
	return mutations
}
