package patient_diseas

import (
	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	pb_models "github.com/MediStatTech/patient-client/pb/go/models/v1"
)

func patientDiseasPropsToPb(patientDiseas domain.PatientDiseasProps) *pb_models.PatientDiseas_Read {
	return &pb_models.PatientDiseas_Read{
		PatientId: patientDiseas.PatientID,
		DiseasId:  patientDiseas.DiseasID,
	}
}
