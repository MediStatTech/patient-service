package get

import "github.com/MediStatTech/patient-service/internal/app/patient/domain"

type Request struct {
	// Empty for get all
}

type Response struct {
	Patients []domain.PatientProps
}
