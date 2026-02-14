package app

import (
	"fmt"

	"github.com/MediStatTech/patient-service/internal/app/patient"
	"github.com/MediStatTech/patient-service/internal/app/patient/usecases"
	"github.com/MediStatTech/patient-service/pkg"
)

type Facade struct {
	Member *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {
	member, err := member.New(pkg)
	if err != nil {
		return nil, fmt.Errorf("failed to create member: %w", err)
	}

	return &Facade{
		Member: member.UseCases,
	}, nil
}
