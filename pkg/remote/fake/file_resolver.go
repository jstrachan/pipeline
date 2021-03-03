package fake

import (
	"github.com/pkg/errors"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned/scheme"
	"github.com/tektoncd/pipeline/pkg/remote"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"path/filepath"
	"strings"
)

// FileResolver implements the Resolver interface files
type FileResolver struct {
	Dir string
}

// NewResolver is a convenience function to return a new OCI resolver instance as a remote.Resolver with a short, 1m
// timeout for resolving an individual image.
func NewFileResolver(dir string) remote.Resolver {
	return &FileResolver{Dir: dir}
}

// List returns the list of objects
func (r *FileResolver) List() ([]remote.ResolvedObject, error) {
	return nil, nil
}

// Get returns the object for the given kind and name
func (r *FileResolver) Get(kind, name string) (runtime.Object, error) {
	// lets strip any git SHA
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
