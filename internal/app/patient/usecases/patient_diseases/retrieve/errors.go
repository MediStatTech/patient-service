package retrieve

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToRetrievePatientDiseas = errors.NewGRPCError(codes.Internal, "failed to retrieve patient diseas")
	errInvalidRequest                = errors.NewGRPCError(codes.InvalidArgument, "invalid request: patient_id and diseases_id are required")
)
