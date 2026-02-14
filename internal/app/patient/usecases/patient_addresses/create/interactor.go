package create

import (
	"context"
	"time"

	"github.com/MediStatTech/patient-service/internal/app/patient/contracts"
	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	"github.com/MediStatTech/patient-service/pkg/commitplan"
)

type Interactor struct {
	addressesRepo contracts.PatientAddressesRepo
	committer     contracts.Committer
	logger        contracts.Logger
}

func New(
	addressesRepo contracts.PatientAddressesRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		addressesRepo: addressesRepo,
		committer:     committer,
		logger:        logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.PatientID == "" || req.Line1 == "" || req.City == "" || req.State == "" {
		return nil, errInvalidRequest
	}

	now := time.Now().UTC()
	address := domain.NewPatientAddress(
		req.PatientID,
		req.Line1,
		req.City,
		req.State,
		now,
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.addressesRepo.CreateMut(address))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return nil, errFailedToCreatePatientAddress.SetInternal(err)
	}

	return &Response{
		PlaceID: address.PlaceID(),
	}, nil
}
