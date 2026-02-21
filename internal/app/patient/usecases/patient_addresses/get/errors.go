package get

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetPatientAddresses = errors.NewGRPCError(codes.Internal, "failed to get patient addresses")
	errInvalidRequest              = errors.NewGRPCError(codes.InvalidArgument, "invalid request: patient_id is required")
)
