package create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCreatePatientContactInfo = errors.NewGRPCError(codes.Internal, "failed to create patient contact info")
	errInvalidRequest                   = errors.NewGRPCError(codes.InvalidArgument, "invalid request: patient_id, phone, and email are required")
)
