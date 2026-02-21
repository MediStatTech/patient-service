package usecases

import (
	patient_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patients/create"
	patient_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patients/get"
	patient_get_by_staff_id "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patients/get_by_staff_id"
	patient_retrieve "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patients/retrieve"
	patient_address_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_addresses/create"
	patient_address_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_addresses/get"
	patient_contact_info_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_contact_infos/create"
	patient_contact_info_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_contact_infos/get"
	patient_contact_info_retrieve "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_contact_infos/retrieve"
	patient_diseas_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_diseases/create"
	patient_diseas_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_diseases/get"
	patient_diseas_retrieve "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_diseases/retrieve"
	"github.com/MediStatTech/patient-service/internal/app/patient/usecases/uc_options"
)

type Facade struct {
	PatientCreate       *patient_create.Interactor
	PatientGet          *patient_get.Interactor
	PatientGetByStaffID *patient_get_by_staff_id.Interactor
	PatientRetrieve     *patient_retrieve.Interactor

	PatientAddressCreate *patient_address_create.Interactor
	PatientAddressGet    *patient_address_get.Interactor

	PatientContactInfoCreate   *patient_contact_info_create.Interactor
	PatientContactInfoGet      *patient_contact_info_get.Interactor
	PatientContactInfoRetrieve *patient_contact_info_retrieve.Interactor

	PatientDiseasCreate   *patient_diseas_create.Interactor
	PatientDiseasGet      *patient_diseas_get.Interactor
	PatientDiseasRetrieve *patient_diseas_retrieve.Interactor
}

func New(o *uc_options.Options) *Facade {
	return &Facade{
		PatientCreate:       patient_create.New(o.PatientsRepo, o.Committer, o.Logger),
		PatientGet:          patient_get.New(o.PatientsRepo, o.Logger),
		PatientGetByStaffID: patient_get_by_staff_id.New(o.PatientsRepo, o.Logger),
		PatientRetrieve:     patient_retrieve.New(o.PatientsRepo, o.Logger),

		PatientAddressCreate: patient_address_create.New(o.PatientAddressesRepo, o.Committer, o.Logger),
		PatientAddressGet:    patient_address_get.New(o.PatientAddressesRepo, o.Logger),

		PatientContactInfoCreate:   patient_contact_info_create.New(o.PatientContactInfosRepo, o.Committer, o.Logger),
		PatientContactInfoGet:      patient_contact_info_get.New(o.PatientContactInfosRepo, o.Logger),
		PatientContactInfoRetrieve: patient_contact_info_retrieve.New(o.PatientContactInfosRepo, o.Logger),

		PatientDiseasCreate:   patient_diseas_create.New(o.PatientDiseasesRepo, o.Committer, o.Logger),
		PatientDiseasGet:      patient_diseas_get.New(o.PatientDiseasesRepo, o.Logger),
		PatientDiseasRetrieve: patient_diseas_retrieve.New(o.PatientDiseasesRepo, o.Logger),
	}
}
