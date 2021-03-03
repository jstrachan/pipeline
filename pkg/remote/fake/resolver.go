package fake

import (
	"github.com/tektoncd/pipeline/pkg/remote"
	"k8s.io/apimachinery/pkg/runtime"
)

// Resolver fake implementation of the Resolver interface
type Resolver struct {
	ResolvedObjects []remote.ResolvedObject
	Objects         map[string]runtime.Object
	Error           error
}

// NewResolver is a convenience function to return a new OCI resolver instance as a remote.Resolver with a short, 1m
// timeout for resolving an individual image.
func NewResolver(resolvedObjects []remote.ResolvedObject, objects map[string]runtime.Object) remote.Resolver {
	if objects == nil {
		objects = map[string]runtime.Object{}
	}
	return &Resolver{ResolvedObjects: resolvedObjects, Objects: objects}
}

// List returns the list of objects
func (r *Resolver) List() ([]remote.ResolvedObject, error) {
	if r.Error != nil {
		return nil, r.Error
	}
	return r.ResolvedObjects, nil
}

// Get returns the object for the given kind and name
func (r *Resolver) Get(kind, name string) (runtime.Object, error) {
	if r.Error != nil {
		return nil, r.Error
	}
	return r.Objects[name], nil
}
