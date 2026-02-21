package patient

import (
	"context"
	"time"

	patient_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patients/create"
	patient_get_by_staff_id "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patients/get_by_staff_id"
	patient_retrieve "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patients/retrieve"
	get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patients/get"
	pb_models "github.com/MediStatTech/patient-client/pb/go/models/v1"
	pb_services "github.com/MediStatTech/patient-client/pb/go/services/v1"
)

func (h *Handler) PatientCreate(
	ctx context.Context,
	req *pb_services.PatientCreateRequest,
) (*pb_services.PatientCreateReply, error) {
	if req == nil || req.Patient == nil {
		return nil, errRequestNil
	}

	patientData := req.Patient
	if patientData.FirstName == "" || patientData.LastName == "" || patientData.Gender == "" || patientData.Dob == "" {
		return nil, errInvalidPatientData
	}

	dob, err := time.Parse("2006-01-02", patientData.Dob)
	if err != nil {
		return nil, errInvalidPatientData
	}

	var staffID *string
	if patientData.StaffId != nil {
		staffID = patientData.StaffId
	}

	resp, err := h.commands.PatientCreate.Execute(ctx, patient_create.Request{
		FirstName: patientData.FirstName,
		LastName:  patientData.LastName,
		Gender:    patientData.Gender,
		Dob:       dob,
		StaffID:   staffID,
	})
	if err != nil {
		h.pkg.Logger.Error("Failed to create patient", map[string]any{"error": err})
		return nil, err
	}

	retrieveResp, err := h.queries.PatientRetrieve.Execute(ctx, patient_retrieve.Request{
		PatientID: resp.PatientID,
	})
	if err != nil {
		h.pkg.Logger.Error("Failed to get created patient", map[string]any{"error": err})
		return nil, err
	}

	return &pb_services.PatientCreateReply{
		Patient: patientPropsToPb(retrieveResp.Patient),
	}, nil
}

func (h *Handler) PatientGet(
	ctx context.Context,
	req *pb_services.PatientGetRequest,
) (*pb_services.PatientGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.PatientGet.Execute(ctx, get.Request{})
	if err != nil {
		h.pkg.Logger.Error("Failed to get patients", map[string]any{"error": err})
		return nil, err
	}

	if len(resp.Patients) == 0 {
		return &pb_services.PatientGetReply{
			Patients: []*pb_models.Patient_Read{},
		}, nil
	}

	patients := make([]*pb_models.Patient_Read, 0, len(resp.Patients))
	for _, patient := range resp.Patients {
		patients = append(patients, patientPropsToPb(patient))
	}

	return &pb_services.PatientGetReply{
		Patients: patients,
	}, nil
}

func (h *Handler) PatientGetByStaffID(
	ctx context.Context,
	req *pb_services.PatientGetByStaffIDRequest,
) (*pb_services.PatientPatientGetByStaffIDReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	if req.StaffId == "" {
		return nil, errInvalidPatientData
	}

	resp, err := h.queries.PatientGetByStaffID.Execute(ctx, patient_get_by_staff_id.Request{
		StaffID: req.StaffId,
	})
	if err != nil {
		h.pkg.Logger.Error("Failed to get patients by staff id", map[string]any{"error": err})
		return nil, err
	}

	if len(resp.Patients) == 0 {
		return &pb_services.PatientPatientGetByStaffIDReply{
			Patients: []*pb_models.Patient_Read{},
		}, nil
	}

	patients := make([]*pb_models.Patient_Read, 0, len(resp.Patients))
	for _, patient := range resp.Patients {
		patients = append(patients, patientPropsToPb(patient))
	}

	return &pb_services.PatientPatientGetByStaffIDReply{
		Patients: patients,
	}, nil
}

func (h *Handler) PatientRetrieve(
	ctx context.Context,
	req *pb_services.PatientRetrieveRequest,
) (*pb_services.PatientRetrieveReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	if req.PatientId == "" {
		return nil, errInvalidPatientData
	}

	resp, err := h.queries.PatientRetrieve.Execute(ctx, patient_retrieve.Request{
		PatientID: req.PatientId,
	})
	if err != nil {
		h.pkg.Logger.Error("Failed to retrieve patient", map[string]any{"error": err})
		return nil, err
	}

	return &pb_services.PatientRetrieveReply{
		Patient: patientPropsToPb(resp.Patient),
	}, nil
}
