package get_by_staff_id

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetPatientsByStaffID = errors.NewGRPCError(codes.Internal, "failed to get patients by staff id")
	errInvalidRequest               = errors.NewGRPCError(codes.InvalidArgument, "invalid request: staff_id is required")
)
