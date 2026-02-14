package get

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetPatientDiseases = errors.NewGRPCError(codes.Internal, "failed to get patient diseases")
)
