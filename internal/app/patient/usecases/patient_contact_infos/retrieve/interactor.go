package retrieve

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
	if req.PatientID == "" {
		return nil, errInvalidRequest
	}

	contactInfos, err := it.contactInfosRepo.FindByPatientIDAndPrimary(ctx, req.PatientID)
	if err != nil {
		return nil, errFailedToRetrievePatientContactInfo.SetInternal(err)
	}

	if contactInfos == nil {
		return nil, errNotFound
	}

	return &Response{
		PatientContactInfo: *contactInfos,
	}, nil
}
