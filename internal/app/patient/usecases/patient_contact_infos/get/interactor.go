package get

import (
	"context"

	"github.com/MediStatTech/patient-service/internal/app/patient/contracts"
)

type Interactor struct {
	contactInfosRepo contracts.PatientContactInfosRepo
	logger           contracts.Logger
}

func New(
	contactInfosRepo contracts.PatientContactInfosRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		contactInfosRepo: contactInfosRepo,
		logger:           logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	contactInfos, err := it.contactInfosRepo.FindAll(ctx)
	if err != nil {
		return nil, errFailedToGetPatientContactInfos.SetInternal(err)
	}

	return &Response{
		PatientContactInfos: contactInfos,
	}, nil
}
