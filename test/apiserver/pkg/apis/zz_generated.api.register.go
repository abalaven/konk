// Code generated by apiregister-gen. DO NOT EDIT.

package apis

import (
	"github.com/infobloxopen/konk/test/apiserver/pkg/apis/example"
	_ "github.com/infobloxopen/konk/test/apiserver/pkg/apis/example/install" // Install the example group
	examplev1alpha1 "github.com/infobloxopen/konk/test/apiserver/pkg/apis/example/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

var (
	localSchemeBuilder = runtime.SchemeBuilder{
		examplev1alpha1.AddToScheme,
	}
	AddToScheme = localSchemeBuilder.AddToScheme
)

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetExampleAPIBuilder(),
	}
}

var exampleApiGroup = builders.NewApiGroupBuilder(
	"example.infoblox.com",
	"github.com/infobloxopen/konk/test/apiserver/pkg/apis/example").
	WithUnVersionedApi(example.ApiVersion).
	WithVersionedApis(
		examplev1alpha1.ApiVersion,
	).
	WithRootScopedKinds()

func GetExampleAPIBuilder() *builders.APIGroupBuilder {
	return exampleApiGroup
}
