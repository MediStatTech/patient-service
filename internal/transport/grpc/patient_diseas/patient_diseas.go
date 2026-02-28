package patient_diseas

import (
	"context"

	pb_models "github.com/MediStatTech/patient-client/pb/go/models/v1"
	pb_services "github.com/MediStatTech/patient-client/pb/go/services/v1"
	patient_diseas_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_diseases/create"
	patient_diseas_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_diseases/get"
	patient_diseas_retrieve "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_diseases/retrieve"
)

func (h *Handler) PatientDiseasCreate(
	ctx context.Context,
	req *pb_services.PatientDiseasCreateRequest,
) (*pb_services.PatientDiseasCreateReply, error) {
	if req == nil || req.PatientDiseas == nil {
		return nil, errRequestNil
	}

	patientDiseasData := req.PatientDiseas
	if patientDiseasData.PatientId == "" || patientDiseasData.DiseasId == "" {
		return nil, errInvalidPatientDiseasData
	}

	resp, err := h.commands.PatientDiseasCreate.Execute(ctx, patient_diseas_create.Request{
		PatientID: patientDiseasData.PatientId,
		DiseasID:  patientDiseasData.DiseasId,
	})
	if err != nil {
		h.pkg.Logger.Error("Failed to create patient diseas", map[string]any{"error": err})
		return nil, err
	}

	diseasesResp, err := h.queries.PatientDiseasGet.Execute(ctx, patient_diseas_get.Request{
		PatientID: req.PatientDiseas.PatientId,
	})
	if err != nil {
		h.pkg.Logger.Error("Failed to get created patient diseas", map[string]any{"error": err})
		return nil, err
	}

	var createdDiseas *pb_models.PatientDiseas_Read
	for _, pd := range diseasesResp.PatientDiseases {
		if pd.PatientID == resp.PatientID && pd.DiseasID == resp.DiseasID {
			createdDiseas = patientDiseasPropsToPb(pd)
			break
		}
	}

	return &pb_services.PatientDiseasCreateReply{
		PatientDiseas: createdDiseas,
	}, nil
}

func (h *Handler) PatientDiseasGet(
	ctx context.Context,
	req *pb_services.PatientDiseasGetRequest,
) (*pb_services.PatientDiseasGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.PatientDiseasGet.Execute(ctx, patient_diseas_get.Request{
		PatientID: req.PatientId,
	})
	if err != nil {
		h.pkg.Logger.Error("Failed to get patient diseases", map[string]any{"error": err})
		return nil, err
	}

	if len(resp.PatientDiseases) == 0 {
		return &pb_services.PatientDiseasGetReply{
			PatientDiseases: []*pb_models.PatientDiseas_Read{},
		}, nil
	}

	diseases := make([]*pb_models.PatientDiseas_Read, 0, len(resp.PatientDiseases))
	for _, patientDiseas := range resp.PatientDiseases {
		diseases = append(diseases, patientDiseasPropsToPb(patientDiseas))
	}

	return &pb_services.PatientDiseasGetReply{
		PatientDiseases: diseases,
	}, nil
}

func (h *Handler) PatientDiseasRetrieve(
	ctx context.Context,
	req *pb_services.PatientDiseasRetrieveRequest,
) (*pb_services.PatientDiseasRetrieveReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.PatientDiseasRetrieve.Execute(ctx, patient_diseas_retrieve.Request{
		PatientID: req.PatientId,
		DiseasID:  req.DiseasesId,
	})
	if err != nil {
		h.pkg.Logger.Error("Failed to retrieve patient diseas", map[string]any{"error": err})
		return nil, err
	}

	return &pb_services.PatientDiseasRetrieveReply{
		PatientDisease: patientDiseasPropsToPb(resp.PatientDiseas),
	}, nil
}
