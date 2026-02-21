package retrieve

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
	if req.PatientID == "" {
		return nil, errInvalidRequest
	}

	patient, err := it.patientsRepo.FindByID(ctx, req.PatientID)
	if err != nil {
		return nil, errFailedToRetrievePatient.SetInternal(err)
	}

	return &Response{
		Patient: patient,
	}, nil
}
