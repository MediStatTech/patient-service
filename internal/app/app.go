package app

import (
	"fmt"

	"github.com/MediStatTech/patient-service/internal/app/patient"
	"github.com/MediStatTech/patient-service/internal/app/patient/usecases"
	"github.com/MediStatTech/patient-service/pkg"
)

type Facade struct {
	Patient *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {
	patientFacade, err := patient.New(pkg)
	if err != nil {
		return nil, fmt.Errorf("failed to create patient: %w", err)
	}

	return &Facade{
		Patient: patientFacade.UseCases,
	}, nil
}
