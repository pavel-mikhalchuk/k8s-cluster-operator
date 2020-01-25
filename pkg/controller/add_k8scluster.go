package controller

import (
	"github.com/pavel-mikhalchuk/k8s-cluster-operator/pkg/controller/k8scluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, k8scluster.Add)
}
