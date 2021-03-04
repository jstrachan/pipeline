package cache

import (
	"github.com/pkg/errors"
	"github.com/tektoncd/pipeline/pkg/remote"
	"k8s.io/apimachinery/pkg/runtime"
)

// Resolver fake implementation of the Resolver interface
type Resolver struct {
	delegate remote.Resolver
	cache    map[string]runtime.Object
	list     []remote.ResolvedObject
	listed   bool
}

// NewResolver caches results so that we don't try to resolve the same resources multiple times
func NewResolver(delegate remote.Resolver) remote.Resolver {
	return &Resolver{
		delegate: delegate,
		cache:    map[string]runtime.Object{},
	}
}

// List returns the list of objects
func (r *Resolver) List() ([]remote.ResolvedObject, error) {
	if !r.listed {
		var err error
		r.list, err = r.delegate.List()
		if err != nil {
			return r.list, errors.Wrapf(err, "failed to ")
		}
		r.listed = true
	}
	return r.list, nil
}

// Get returns the object for the given kind and name
func (r *Resolver) Get(kind, name string) (runtime.Object, error) {
	key := kind + "/" + name

	obj, ok := r.cache[key]
	if ok {
		return obj, nil
	}

	obj, err := r.delegate.Get(kind, name)
	if err != nil {
		return obj, err
	}
	r.cache[key] = obj
	return obj, nil
}
