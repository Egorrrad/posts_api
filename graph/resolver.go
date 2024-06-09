package graph

import (
	"GraphQL_api/internal"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Store internal.DataStorage
}
