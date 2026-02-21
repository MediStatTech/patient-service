package get

import (
	"context"

	"github.com/MediStatTech/patient-service/internal/app/patient/contracts"
)

type Interactor struct {
	addressesRepo contracts.PatientAddressesRepo
	logger        contracts.Logger
}

func New(
	addressesRepo contracts.PatientAddressesRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		addressesRepo: addressesRepo,
		logger:        logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.PatientID == "" {
		return nil, errInvalidRequest
	}

	addresses, err := it.addressesRepo.FindByPatientID(ctx, req.PatientID)
	if err != nil {
		return nil, errFailedToGetPatientAddresses.SetInternal(err)
	}

	return &Response{
		PatientAddresses: addresses,
	}, nil
}
