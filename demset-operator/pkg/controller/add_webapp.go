package controller

import (
	"github.com/DevOpsLK/demset-operator/pkg/controller/webapp"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, webapp.Add)
}
