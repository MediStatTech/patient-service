package patient

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errRequestNil        = errors.NewGRPCError(codes.InvalidArgument, "request cannot be nil")
	errInvalidPatientData = errors.NewGRPCError(codes.InvalidArgument, "invalid patient data")
)
