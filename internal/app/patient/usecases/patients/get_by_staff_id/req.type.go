package get_by_staff_id

import "github.com/MediStatTech/patient-service/internal/app/patient/domain"

type Request struct {
	StaffID string
}

type Response struct {
	Patients []domain.PatientProps
}
