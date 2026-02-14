package create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCreatePatientDiseas = errors.NewGRPCError(codes.Internal, "failed to create patient diseas")
	errInvalidRequest              = errors.NewGRPCError(codes.InvalidArgument, "invalid request: patient_id and diseas_id are required")
)
