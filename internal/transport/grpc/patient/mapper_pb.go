package patient

import (
	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	pb_models "github.com/MediStatTech/patient-client/pb/go/models/v1"
)

func patientPropsToPb(patient domain.PatientProps) *pb_models.Patient_Read {
	pb := &pb_models.Patient_Read{
		PatientId: patient.PatientID,
		FirstName: patient.FirstName,
		LastName:  patient.LastName,
		Gender:    patient.Gender,
		Dob:       patient.Dob.Format("2006-01-02"),
	}
	
	if patient.StaffID != nil {
		pb.StaffId = patient.StaffID
	}
	
	return pb
}
