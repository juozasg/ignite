package v1alpha1

import (
	"github.com/weaveworks/ignite/pkg/apis/ignite"
	"k8s.io/apimachinery/pkg/conversion"
)

// Convert_ignite_VMSpec_To_v1alpha1_VMSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_ignite_VMSpec_To_v1alpha1_VMSpec(in *ignite.VMSpec, out *VMSpec, s conversion.Scope) error {
	// VMSpecStorage are not supported by v1alpha1, so just ignore the warning by calling this manually
	return autoConvert_ignite_VMSpec_To_v1alpha1_VMSpec(in, out, s)
}

// Convert_ignite_VMSpec_To_v1alpha1_VMSpec calls the autogenerated conversion function along with custom conversion logic
func Convert_v1alpha1_VMSpec_To_ignite_VMSpec(in *VMSpec, out *ignite.VMSpec, s conversion.Scope) error {
	// VMSpecStorage is not supported by v1alpha1, so just ignore the warning by calling this manually
	return autoConvert_v1alpha1_VMSpec_To_ignite_VMSpec(in, out, s)
}
