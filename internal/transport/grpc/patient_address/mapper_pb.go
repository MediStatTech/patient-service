package patient_address

import (
	"github.com/MediStatTech/patient-service/internal/app/patient/domain"
	pb_models "github.com/MediStatTech/patient-client/pb/go/models/v1"
)

func patientAddressPropsToPb(address domain.PatientAddressProps) *pb_models.PatientAddress_Read {
	return &pb_models.PatientAddress_Read{
		PatientId: address.PatientID,
		PlaceId:   address.PlaceID,
		Line_1:    address.Line1,
		City:      address.City,
		State:     address.State,
	}
}
