package patient

import (
	patient_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patients/create"
	patient_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patients/get"
	s_options "github.com/MediStatTech/patient-service/internal/app/options"
	"github.com/MediStatTech/patient-service/pkg"
)

type Handler struct {
	pkg      *pkg.Facade
	commands *Commands
	queries  *Queries
}

type Commands struct {
	PatientCreate *patient_create.Interactor
}

type Queries struct {
	PatientGet *patient_get.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		commands: &Commands{
			PatientCreate: opts.App.Patient.PatientCreate,
		},
		queries: &Queries{
			PatientGet: opts.App.Patient.PatientGet,
		},
	}
}
