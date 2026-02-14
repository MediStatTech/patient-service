package patient_contact_info

import (
	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	pb_models "github.com/MediStatTech/patient-client/pb/go/models/v1"
)

func patientContactInfoPropsToPb(contactInfo domain.PatientContactInfoProps) *pb_models.PatientContactInfo_Read {
	return &pb_models.PatientContactInfo_Read{
		PatientId: contactInfo.PatientID,
		ContactId: contactInfo.ContactID,
		Phone:     contactInfo.Phone,
		Email:     contactInfo.Email,
		Primary:   contactInfo.Primary,
	}
}
