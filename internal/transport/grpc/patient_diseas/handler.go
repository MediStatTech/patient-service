package patient_diseas

import (
	s_options "github.com/MediStatTech/patient-service/internal/app/options"
	patient_diseas_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_diseases/create"
	patient_diseas_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_diseases/get"
	patient_diseas_retrieve "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_diseases/retrieve"
	"github.com/MediStatTech/patient-service/pkg"
)

type Handler struct {
	pkg      *pkg.Facade
	commands *Commands
	queries  *Queries
}

type Commands struct {
	PatientDiseasCreate *patient_diseas_create.Interactor
}

type Queries struct {
	PatientDiseasGet      *patient_diseas_get.Interactor
	PatientDiseasRetrieve *patient_diseas_retrieve.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		commands: &Commands{
			PatientDiseasCreate: opts.App.Patient.PatientDiseasCreate,
		},
		queries: &Queries{
			PatientDiseasGet:      opts.App.Patient.PatientDiseasGet,
			PatientDiseasRetrieve: opts.App.Patient.PatientDiseasRetrieve,
		},
	}
}
