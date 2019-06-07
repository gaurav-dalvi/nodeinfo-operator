package controller

import (
	"github.com/gaurav-dalvi/nodeinfo-operator/pkg/controller/nodeinfo"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nodeinfo.Add)
}
