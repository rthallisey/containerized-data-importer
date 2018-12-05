package operator

import (
	"kubevirt.io/containerized-data-importer/pkg/apis/cdi/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)

	AddToManagerFuncs = append(AddToManagerFuncs, Add)
}
