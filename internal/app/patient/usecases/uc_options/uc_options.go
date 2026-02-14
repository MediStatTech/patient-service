package uc_options

import (
	"github.com/MediStatTech/patient-service/internal/app/patient/contracts"
)

type Options struct {
	Committer contracts.Committer
	Logger    contracts.Logger
}
