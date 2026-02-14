package patient_contact_info

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errRequestNil                   = errors.NewGRPCError(codes.InvalidArgument, "request cannot be nil")
	errInvalidPatientContactInfoData = errors.NewGRPCError(codes.InvalidArgument, "invalid patient contact info data")
)
