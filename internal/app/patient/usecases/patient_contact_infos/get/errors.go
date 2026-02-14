package get

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToGetPatientContactInfos = errors.NewGRPCError(codes.Internal, "failed to get patient contact infos")
)
