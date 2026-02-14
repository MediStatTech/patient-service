package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
	"github.com/google/uuid"
)

// ============================================================================
// Patient Addresses Repository
// ============================================================================

type PatientAddressesRepository struct {
	queries *Queries
}

func NewPatientAddressesRepository(db *sql.DB) *PatientAddressesRepository {
	return &PatientAddressesRepository{
		queries: New(db),
	}
}

func (r *PatientAddressesRepository) FindAll(ctx context.Context) ([]domain.PatientAddressProps, error) {
	addresses, err := r.queries.ListPatientAddresses(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PatientAddressProps, 0, len(addresses))
	for _, address := range addresses {
		result = append(result, toPatientAddressProps(address))
	}

	return result, nil
}

func (r *PatientAddressesRepository) FindByID(ctx context.Context, placeID string) (domain.PatientAddressProps, error) {
	id, err := uuid.Parse(placeID)
	if err != nil {
		return domain.PatientAddressProps{}, err
	}

	address, err := r.queries.GetPatientAddress(ctx, id)
	if err != nil {
		return domain.PatientAddressProps{}, err
	}

	return toPatientAddressProps(address), nil
}

func (r *PatientAddressesRepository) FindByPatientID(ctx context.Context, patientID string) ([]domain.PatientAddressProps, error) {
	id, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	addresses, err := r.queries.ListPatientAddressesByPatientID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PatientAddressProps, 0, len(addresses))
	for _, address := range addresses {
		result = append(result, toPatientAddressProps(address))
	}

	return result, nil
}

func (r *PatientAddressesRepository) CreateMut(address *domain.PatientAddress) *postgres.Mutation {
	return postgres.NewMutation(
		CreatePatientAddress,
		patientAddressToCreateParams(address)...,
	)
}

func (r *PatientAddressesRepository) UpdateMut(address *domain.PatientAddress) *postgres.Mutation {
	return postgres.NewMutation(
		UpdatePatientAddress,
		patientAddressToUpdateParams(address)...,
	)
}

func (r *PatientAddressesRepository) DeleteMut(placeID string) *postgres.Mutation {
	id, _ := uuid.Parse(placeID)
	return postgres.NewMutation(
		DeletePatientAddress,
		id,
	)
}

func (r *PatientAddressesRepository) CreateBatchMut(addresses []*domain.PatientAddress) []*postgres.Mutation {
	mutations := make([]*postgres.Mutation, 0, len(addresses))
	for _, address := range addresses {
		mutations = append(mutations, r.CreateMut(address))
	}
	return mutations
}
