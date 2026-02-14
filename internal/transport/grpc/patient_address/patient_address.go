package patient_address

import (
	"context"

	patient_address_create "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_addresses/create"
	patient_address_get "github.com/MediStatTech/patient-service/internal/app/patient/usecases/patient_addresses/get"
	pb_services "github.com/MediStatTech/patient-client/pb/go/services/v1"
	pb_models "github.com/MediStatTech/patient-client/pb/go/models/v1"
)

func (h *Handler) PatientAddressCreate(
	ctx context.Context,
	req *pb_services.PatientAddressCreateRequest,
) (*pb_services.PatientAddressCreateReply, error) {
	if req == nil || req.PatientAddress == nil {
		return nil, errRequestNil
	}

	addressData := req.PatientAddress
	if addressData.PatientId == "" || addressData.Line_1 == "" || addressData.City == "" || addressData.State == "" {
		return nil, errInvalidPatientAddressData
	}

	resp, err := h.commands.PatientAddressCreate.Execute(ctx, patient_address_create.Request{
		PatientID: addressData.PatientId,
		Line1:     addressData.Line_1,
		City:      addressData.City,
		State:     addressData.State,
	})
	if err != nil {
		h.pkg.Logger.Error("Failed to create patient address", map[string]any{"error": err})
		return nil, err
	}

	addressesResp, err := h.queries.PatientAddressGet.Execute(ctx, patient_address_get.Request{})
	if err != nil {
		h.pkg.Logger.Error("Failed to get created patient address", map[string]any{"error": err})
		return nil, err
	}

	var createdAddress *pb_models.PatientAddress_Read
	for _, a := range addressesResp.PatientAddresses {
		if a.PlaceID == resp.PlaceID {
			createdAddress = patientAddressPropsToPb(a)
			break
		}
	}

	return &pb_services.PatientAddressCreateReply{
		PatientAddress: createdAddress,
	}, nil
}

func (h *Handler) PatientAddressGet(
	ctx context.Context,
	req *pb_services.PatientAddressGetRequest,
) (*pb_services.PatientAddressGetReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.PatientAddressGet.Execute(ctx, patient_address_get.Request{})
	if err != nil {
		h.pkg.Logger.Error("Failed to get patient addresses", map[string]any{"error": err})
		return nil, err
	}

	if len(resp.PatientAddresses) == 0 {
		return &pb_services.PatientAddressGetReply{
			PatientAddresses: []*pb_models.PatientAddress_Read{},
		}, nil
	}

	addresses := make([]*pb_models.PatientAddress_Read, 0, len(resp.PatientAddresses))
	for _, address := range resp.PatientAddresses {
		addresses = append(addresses, patientAddressPropsToPb(address))
	}

	return &pb_services.PatientAddressGetReply{
		PatientAddresses: addresses,
	}, nil
}
