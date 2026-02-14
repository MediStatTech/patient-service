package patient_address

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errRequestNil              = errors.NewGRPCError(codes.InvalidArgument, "request cannot be nil")
	errInvalidPatientAddressData = errors.NewGRPCError(codes.InvalidArgument, "invalid patient address data")
)
