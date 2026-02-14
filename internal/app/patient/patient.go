package member

import (
	"github.com/MediStatTech/patient-service/internal/app/patient/usecases"
	"github.com/MediStatTech/patient-service/internal/app/patient/usecases/uc_options"
	"github.com/MediStatTech/patient-service/pkg"
)

type Facade struct {
	pkg      *pkg.Facade
	UseCases *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {

	useCasesInstance := usecases.New(&uc_options.Options{
		Committer: pkg.Committer,
		Logger:    pkg.Logger,
	})

	return &Facade{
		pkg:      pkg,
		UseCases: useCasesInstance,
	}, nil
}
