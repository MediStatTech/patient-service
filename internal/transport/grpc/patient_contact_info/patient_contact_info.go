package patient_contact_info

import (
	"context"

	patient_contact_info_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_contact_infos/create"
	patient_contact_info_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_contact_infos/get"
	pb_services "github.com/MediStatTech/patient-client/pb/go/services/v1"
	pb_models "github.com/MediStatTech/patient-client/pb/go/models/v1"
)

func (h *Handler) PatientContactInfoCreate(
	ctx context.Context,
	req *pb_services.PatientContactInfoCreateRequest,
) (*pb_services.PatientContactInfoCreateReply, error) {
	if req == nil || req.PatientContactInfo == nil {
		return nil, errRequestNil
	}

	contactInfoData := req.PatientContactInfo
	if contactInfoData.PatientId == "" || contactInfoData.Phone == "" || contactInfoData.Email == "" {
		return nil, errInvalidPatientContactInfoData
	}

	resp, err := h.commands.PatientContactInfoCreate.Execute(ctx, patient_contact_info_create.Request{
		PatientID: contactInfoData.PatientId,
		Phone:     contactInfoData.Phone,
		Email:     contactInfoData.Email,
		Primary:   contactInfoData.Primary,
	})
	if err != nil {
		h.pkg.Logger.Error("Failed to create patient contact info", map[string]any{"error": err})
		return nil, err
	}

	contactInfosResp, err := h.queries.PatientContactInfoGet.Execute(ctx, patient_contact_info_get.Request{})
	if err != nil {
		h.pkg.Logger.Error("Failed to get created patient contact info", map[string]any{"error": err})
		return nil, err
	}

	var createdContactInfo *pb_models.PatientContactInfo_Read
	for _, ci := range contactInfosResp.PatientContactInfos {
		if ci.ContactID == resp.ContactID {
			createdContactInfo = patientContactInfoPropsToPb(ci)
			break
		}
	}

	return &pb_services.PatientContactInfoCreateReply{
		PatientContactInfo: createdContactInfo,
	}, nil
}

func (h *Handler) PatientContactInfoGet(
	ctx context.Context,
	req *pb_services.PatientContactInfoGetRequest,
) (*pb_services.PatientContactInfoGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.PatientContactInfoGet.Execute(ctx, patient_contact_info_get.Request{})
	if err != nil {
		h.pkg.Logger.Error("Failed to get patient contact infos", map[string]any{"error": err})
		return nil, err
	}

	if len(resp.PatientContactInfos) == 0 {
		return &pb_services.PatientContactInfoGetReply{
			PatientContactInfos: []*pb_models.PatientContactInfo_Read{},
		}, nil
	}

	contactInfos := make([]*pb_models.PatientContactInfo_Read, 0, len(resp.PatientContactInfos))
	for _, contactInfo := range resp.PatientContactInfos {
		contactInfos = append(contactInfos, patientContactInfoPropsToPb(contactInfo))
	}

	return &pb_services.PatientContactInfoGetReply{
		PatientContactInfos: contactInfos,
	}, nil
}
