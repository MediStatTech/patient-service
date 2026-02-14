package options

import (
	"github.com/MediStatTech/patient-service/internal/app"
	"github.com/MediStatTech/patient-service/pkg"
)

type Options struct {
	App *app.Facade
	PKG *pkg.Facade
}
