package controller

import (
	"github.com/leg100/stok/pkg/controller/workspace"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, workspace.Add)
}
