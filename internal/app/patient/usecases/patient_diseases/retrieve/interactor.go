package retrieve

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
	if req.PatientID == "" || req.DiseasID == "" {
		return nil, errInvalidRequest
	}

	patientDiseas, err := it.diseasesRepo.FindByPatientAndDiseas(ctx, req.PatientID, req.DiseasID)
	if err != nil {
		return nil, errFailedToRetrievePatientDiseas.SetInternal(err)
	}

	return &Response{
		PatientDiseas: patientDiseas,
	}, nil
}
