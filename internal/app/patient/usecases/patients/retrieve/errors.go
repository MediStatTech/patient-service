package retrieve

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToRetrievePatient = errors.NewGRPCError(codes.Internal, "failed to retrieve patient")
	errInvalidRequest          = errors.NewGRPCError(codes.InvalidArgument, "invalid request: patient_id is required")
)
