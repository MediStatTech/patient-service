package create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCreatePatientAddress = errors.NewGRPCError(codes.Internal, "failed to create patient address")
	errInvalidRequest               = errors.NewGRPCError(codes.InvalidArgument, "invalid request: patient_id, line_1, city, and state are required")
)
