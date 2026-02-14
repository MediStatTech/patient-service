package get

import (
	"context"

	"github.com/MediStatTech/patient-service/internal/app/patient/contracts"
)

type Interactor struct {
	patientsRepo contracts.PatientsRepo
	logger       contracts.Logger
}

func New(
	patientsRepo contracts.PatientsRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		patientsRepo: patientsRepo,
		logger:       logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	patients, err := it.patientsRepo.FindAll(ctx)
	if err != nil {
		return nil, errFailedToGetPatients.SetInternal(err)
	}

	return &Response{
		Patients: patients,
	}, nil
}
