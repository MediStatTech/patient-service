package get_by_staff_id

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
	if req.StaffID == "" {
		return nil, errInvalidRequest
	}

	patients, err := it.patientsRepo.FindByStaffID(ctx, req.StaffID)
	if err != nil {
		return nil, errFailedToGetPatientsByStaffID.SetInternal(err)
	}

	return &Response{
		Patients: patients,
	}, nil
}
