package graph

import (
	"github.com/rupam_joshi/star_wars/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.StarWarsService
}
