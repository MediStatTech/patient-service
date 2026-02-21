package retrieve

import "github.com/MediStatTech/patient-service/internal/app/patient/domain"

type Request struct {
	PatientID string
}

type Response struct {
	Patient domain.PatientProps
}
