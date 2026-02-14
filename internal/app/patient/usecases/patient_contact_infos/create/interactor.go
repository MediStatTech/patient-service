package create

import (
	"context"
	"time"

	"github.com/MediStatTech/patient-service/internal/app/patient/contracts"
	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	"github.com/MediStatTech/patient-service/pkg/commitplan"
)

type Interactor struct {
	contactInfosRepo contracts.PatientContactInfosRepo
	committer        contracts.Committer
	logger           contracts.Logger
}

func New(
	contactInfosRepo contracts.PatientContactInfosRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		contactInfosRepo: contactInfosRepo,
		committer:        committer,
		logger:           logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.PatientID == "" || req.Phone == "" || req.Email == "" {
		return nil, errInvalidRequest
	}

	now := time.Now().UTC()
	contactInfo := domain.NewPatientContactInfo(
		req.PatientID,
		req.Phone,
		req.Email,
		req.Primary,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.contactInfosRepo.CreateMut(contactInfo))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreatePatientContactInfo.SetInternal(err)
	}

	return &Response{
		ContactID: contactInfo.ContactID(),
	}, nil
}
