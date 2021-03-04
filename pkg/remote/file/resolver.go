package file

import (
	"github.com/pkg/errors"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned/scheme"
	"github.com/tektoncd/pipeline/pkg/remote"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"path/filepath"
	"strings"
)

// Resolver implements the Resolver interface using files
type Resolver struct {
	Dir string
}

// NewResolver creates a resolver using the given directory
func NewResolver(dir string) remote.Resolver {
	return &Resolver{Dir: dir}
}

// List returns the list of objects
func (r *Resolver) List() ([]remote.ResolvedObject, error) {
	return nil, nil
}

// Get returns the object for the given kind and name
func (r *Resolver) Get(kind, name string) (runtime.Object, error) {
	// lets strip any git SHA so we can be used for testing more easily
	i := strings.LastIndex(name, "@")
	if i > 0 {
		name = name[0:i]
	}
	path := filepath.Join(r.Dir, name)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file %s", path)
	}

	obj, _, err := scheme.Codecs.UniversalDeserializer().Decode(data, nil, nil)
	return obj, err
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal file %s", path)
	}
	return obj, nil
}
