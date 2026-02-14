package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

// ============================================================================
// Patient Diseases Repository
// ============================================================================

type PatientDiseasesRepository struct {
	queries *Queries
}

func NewPatientDiseasesRepository(db *sql.DB) *PatientDiseasesRepository {
	return &PatientDiseasesRepository{
		queries: New(db),
	}
}

func (r *PatientDiseasesRepository) FindAll(ctx context.Context) ([]domain.PatientDiseasProps, error) {
	patientDiseases, err := r.queries.ListPatientDiseases(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PatientDiseasProps, 0, len(patientDiseases))
	for _, pd := range patientDiseases {
		result = append(result, toPatientDiseasProps(pd))
	}

	return result, nil
}

func (r *PatientDiseasesRepository) FindByPatientID(ctx context.Context, patientID string) ([]domain.PatientDiseasProps, error) {
	id, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	patientDiseases, err := r.queries.ListPatientDiseasesByPatientID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PatientDiseasProps, 0, len(patientDiseases))
	for _, pd := range patientDiseases {
		result = append(result, toPatientDiseasProps(pd))
	}

	return result, nil
}

func (r *PatientDiseasesRepository) FindByDiseasID(ctx context.Context, diseasID string) ([]domain.PatientDiseasProps, error) {
	id, err := uuid.Parse(diseasID)
	if err != nil {
		return nil, err
	}

	patientDiseases, err := r.queries.ListPatientDiseasesByDiseasID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PatientDiseasProps, 0, len(patientDiseases))
	for _, pd := range patientDiseases {
		result = append(result, toPatientDiseasProps(pd))
	}

	return result, nil
}

func (r *PatientDiseasesRepository) FindByPatientAndDiseas(ctx context.Context, patientID, diseasID string) (domain.PatientDiseasProps, error) {
	pid, err := uuid.Parse(patientID)
	if err != nil {
		return domain.PatientDiseasProps{}, err
	}

	did, err := uuid.Parse(diseasID)
	if err != nil {
		return domain.PatientDiseasProps{}, err
	}

	pd, err := r.queries.GetPatientDiseas(ctx, GetPatientDiseasParams{
		PatientID: pid,
		DiseasID:  did,
	})
	if err != nil {
		return domain.PatientDiseasProps{}, err
	}

	return toPatientDiseasProps(pd), nil
}

func (r *PatientDiseasesRepository) CreateMut(patientDiseas *domain.PatientDiseas) *postgres.Mutation {
	return postgres.NewMutation(
		CreatePatientDiseas,
		patientDiseasToCreateParams(patientDiseas)...,
	)
}

func (r *PatientDiseasesRepository) UpdateMut(patientDiseas *domain.PatientDiseas) *postgres.Mutation {
	return postgres.NewMutation(
		UpdatePatientDiseas,
		patientDiseasToUpdateParams(patientDiseas)...,
	)
}

func (r *PatientDiseasesRepository) DeleteMut(patientID, diseasID string) *postgres.Mutation {
	pid, _ := uuid.Parse(patientID)
	did, _ := uuid.Parse(diseasID)
	return postgres.NewMutation(
		DeletePatientDiseas,
		DeletePatientDiseasParams{
			PatientID: pid,
			DiseasID:  did,
		},
	)
}

func (r *PatientDiseasesRepository) CreateBatchMut(patientDiseases []*domain.PatientDiseas) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(patientDiseases))
	for _, pd := range patientDiseases {
		mutations = append(mutations, r.CreateMut(pd))
	}
	return mutations
}
