package retrieve

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToRetrievePatientContactInfo = errors.NewGRPCError(codes.Internal, "failed to retrieve patient contact info")
	errInvalidRequest                     = errors.NewGRPCError(codes.InvalidArgument, "invalid request: patient_id is required")
	errNotFound                           = errors.NewGRPCError(codes.NotFound, "patient contact info not found")
)
