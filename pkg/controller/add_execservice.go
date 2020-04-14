package controller

import (
	"github.com/example-inc/exec/pkg/controller/execservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, execservice.Add)
}
