package get

import (
	"context"

	"github.com/MediStatTech/patient-service/internal/app/patient/contracts"
)

type Interactor struct {
	diseasesRepo contracts.PatientDiseasesRepo
	logger       contracts.Logger
}

func New(
	diseasesRepo contracts.PatientDiseasesRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		diseasesRepo: diseasesRepo,
		logger:       logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	diseases, err := it.diseasesRepo.FindAll(ctx)
	if err != nil {
		return nil, errFailedToGetPatientDiseases.SetInternal(err)
	}

	return &Response{
		PatientDiseases: diseases,
	}, nil
}
