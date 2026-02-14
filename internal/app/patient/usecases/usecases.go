package usecases

import (
	"github.com/MediStatTech/patient-service/internal/app/patient/usecases/uc_options"
)

type Facade struct {
}

func New(o *uc_options.Options) *Facade {
	return &Facade{}
}
