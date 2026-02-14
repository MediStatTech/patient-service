package create

import (
	errors "github.com/MediStatTech/MediStat-error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToCreatePatient = errors.NewGRPCError(codes.Internal, "failed to create patient")
	errInvalidRequest        = errors.NewGRPCError(codes.InvalidArgument, "invalid request: first_name, last_name, gender, and dob are required")
)
