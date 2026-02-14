package create

import (
	"context"
	"time"

	"github.com/MediStatTech/patient-service/internal/app/patient/contracts"
	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	"github.com/MediStatTech/patient-service/pkg/commitplan"
)

type Interactor struct {
	diseasesRepo contracts.PatientDiseasesRepo
	committer    contracts.Committer
	logger       contracts.Logger
}

func New(
	diseasesRepo contracts.PatientDiseasesRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		diseasesRepo: diseasesRepo,
		committer:    committer,
		logger:       logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.PatientID == "" || req.DiseasID == "" {
		return nil, errInvalidRequest
	}

	now := time.Now().UTC()
	patientDiseas := domain.NewPatientDiseas(
		req.PatientID,
		req.DiseasID,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.diseasesRepo.CreateMut(patientDiseas))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreatePatientDiseas.SetInternal(err)
	}

	return &Response{
		PatientID: patientDiseas.PatientID(),
		DiseasID:  patientDiseas.DiseasID(),
	}, nil
}
