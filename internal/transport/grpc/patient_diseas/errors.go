package patient_diseas

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errRequestNil            = errors.NewGRPCError(codes.InvalidArgument, "request cannot be nil")
	errInvalidPatientDiseasData = errors.NewGRPCError(codes.InvalidArgument, "invalid patient diseas data")
)
