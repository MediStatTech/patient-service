package patient

import (
	"github.com/MediStatTech/patient-service/internal/app/patient/usecases"
	"github.com/MediStatTech/patient-service/internal/app/patient/usecases/uc_options"
	"github.com/MediStatTech/patient-service/internal/infra/repo"
	"github.com/MediStatTech/patient-service/pkg"
)

type Facade struct {
	pkg      *pkg.Facade
	UseCases *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {
	// Initialize repositories
	patientsRepo := repo.NewPatientsRepository(pkg.Postgres.DB)
	patientAddressesRepo := repo.NewPatientAddressesRepository(pkg.Postgres.DB)
	patientContactInfosRepo := repo.NewPatientContactInfosRepository(pkg.Postgres.DB)
	patientDiseasesRepo := repo.NewPatientDiseasesRepository(pkg.Postgres.DB)

	useCasesInstance := usecases.New(&uc_options.Options{
		Committer:               pkg.Committer,
		Logger:                  pkg.Logger,
		PatientsRepo:            patientsRepo,
		PatientAddressesRepo:    patientAddressesRepo,
		PatientContactInfosRepo: patientContactInfosRepo,
		PatientDiseasesRepo:     patientDiseasesRepo,
	})

	return &Facade{
		pkg:      pkg,
		UseCases: useCasesInstance,
	}, nil
}
