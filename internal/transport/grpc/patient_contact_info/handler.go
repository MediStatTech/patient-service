package patient_contact_info

import (
	s_options "github.com/MediStatTech/patient-service/internal/app/options"
	patient_contact_info_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_contact_infos/create"
	patient_contact_info_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_contact_infos/get"
	patient_contact_info_retrieve "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_contact_infos/retrieve"
	"github.com/MediStatTech/patient-service/pkg"
)

type Handler struct {
	pkg      *pkg.Facade
	commands *Commands
	queries  *Queries
}

type Commands struct {
	PatientContactInfoCreate *patient_contact_info_create.Interactor
}

type Queries struct {
	PatientContactInfoGet      *patient_contact_info_get.Interactor
	PatientContactInfoRetrieve *patient_contact_info_retrieve.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		commands: &Commands{
			PatientContactInfoCreate: opts.App.Patient.PatientContactInfoCreate,
		},
		queries: &Queries{
			PatientContactInfoGet:      opts.App.Patient.PatientContactInfoGet,
			PatientContactInfoRetrieve: opts.App.Patient.PatientContactInfoRetrieve,
		},
	}
}
