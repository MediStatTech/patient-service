package patient_address

import (
	patient_address_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_addresses/create"
	patient_address_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_addresses/get"
	s_options "github.com/MediStatTech/patient-service/internal/app/options"
	"github.com/MediStatTech/patient-service/pkg"
)

type Handler struct {
	pkg      *pkg.Facade
	commands *Commands
	queries  *Queries
}

type Commands struct {
	PatientAddressCreate *patient_address_create.Interactor
}

type Queries struct {
	PatientAddressGet *patient_address_get.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		commands: &Commands{
			PatientAddressCreate: opts.App.Patient.PatientAddressCreate,
		},
		queries: &Queries{
			PatientAddressGet: opts.App.Patient.PatientAddressGet,
		},
	}
}
