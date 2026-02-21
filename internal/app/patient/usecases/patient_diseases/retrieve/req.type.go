package retrieve

import "github.com/MediStatTech/patient-service/internal/app/patient/domain"

type Request struct {
	PatientID string
	DiseasID  string
}

type Response struct {
	PatientDiseas domain.PatientDiseasProps
}
