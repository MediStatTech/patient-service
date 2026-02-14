package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

// ============================================================================
// Patient Contact Infos Repository
// ============================================================================

type PatientContactInfosRepository struct {
	queries *Queries
}

func NewPatientContactInfosRepository(db *sql.DB) *PatientContactInfosRepository {
	return &PatientContactInfosRepository{
		queries: New(db),
	}
}

func (r *PatientContactInfosRepository) FindAll(ctx context.Context) ([]domain.PatientContactInfoProps, error) {
	contactInfos, err := r.queries.ListPatientContactInfos(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PatientContactInfoProps, 0, len(contactInfos))
	for _, contactInfo := range contactInfos {
		result = append(result, toPatientContactInfoProps(contactInfo))
	}

	return result, nil
}

func (r *PatientContactInfosRepository) FindByID(ctx context.Context, contactID string) (domain.PatientContactInfoProps, error) {
	id, err := uuid.Parse(contactID)
	if err != nil {
		return domain.PatientContactInfoProps{}, err
	}

	contactInfo, err := r.queries.GetPatientContactInfo(ctx, id)
	if err != nil {
		return domain.PatientContactInfoProps{}, err
	}

	return toPatientContactInfoProps(contactInfo), nil
}

func (r *PatientContactInfosRepository) FindByPatientID(ctx context.Context, patientID string) ([]domain.PatientContactInfoProps, error) {
	id, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	contactInfos, err := r.queries.ListPatientContactInfosByPatientID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PatientContactInfoProps, 0, len(contactInfos))
	for _, contactInfo := range contactInfos {
		result = append(result, toPatientContactInfoProps(contactInfo))
	}

	return result, nil
}

func (r *PatientContactInfosRepository) CreateMut(contactInfo *domain.PatientContactInfo) *postgres.Mutation {
	return postgres.NewMutation(
		CreatePatientContactInfo,
		patientContactInfoToCreateParams(contactInfo)...,
	)
}

func (r *PatientContactInfosRepository) UpdateMut(contactInfo *domain.PatientContactInfo) *postgres.Mutation {
	return postgres.NewMutation(
		UpdatePatientContactInfo,
		patientContactInfoToUpdateParams(contactInfo)...,
	)
}

func (r *PatientContactInfosRepository) DeleteMut(contactID string) *postgres.Mutation {
	id, _ := uuid.Parse(contactID)
	return postgres.NewMutation(
		DeletePatientContactInfo,
		id,
	)
}

func (r *PatientContactInfosRepository) CreateBatchMut(contactInfos []*domain.PatientContactInfo) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(contactInfos))
	for _, contactInfo := range contactInfos {
		mutations = append(mutations, r.CreateMut(contactInfo))
	}
	return mutations
}
