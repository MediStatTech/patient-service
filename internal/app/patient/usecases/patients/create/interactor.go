package create

import (
	"context"
	"time"

	"github.com/MediStatTech/patient-service/internal/app/patient/contracts"
	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	"github.com/MediStatTech/patient-service/pkg/commitplan"
)

type Interactor struct {
	patientsRepo contracts.PatientsRepo
	committer    contracts.Committer
	logger       contracts.Logger
}

func New(
	patientsRepo contracts.PatientsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		patientsRepo: patientsRepo,
		committer:    committer,
		logger:       logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.FirstName == "" || req.LastName == "" || req.Gender == "" || req.Dob.IsZero() {
		return nil, errInvalidRequest
	}

	now := time.Now().UTC()
	patient := domain.NewPatient(
		req.FirstName,
		req.LastName,
		req.Gender,
		req.Dob,
		req.StaffID,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.patientsRepo.CreateMut(patient))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreatePatient.SetInternal(err)
	}

	return &Response{
		PatientID: patient.PatientID(),
	}, nil
}
