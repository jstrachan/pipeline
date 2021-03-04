package stepper

import (
	"context"
	"fmt"
	"github.com/google/go-containerregistry/pkg/authn/k8schain"
	"github.com/pkg/errors"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/remote"
	"github.com/tektoncd/pipeline/pkg/remote/oci"
	"k8s.io/client-go/kubernetes"
)

type RemoterOptions struct {
	KubeClient        kubernetes.Interface
	Namespace         string
	OCIServiceAccount string
}

func (o *RemoterOptions) CreateRemote(ctx context.Context, uses *v1beta1.Uses) (remote.Resolver, error) {
	if uses.Kind == v1beta1.UsesTypeOCI {
		bundle := uses.Path
		kc, err := k8schain.New(ctx, o.KubeClient, k8schain.Options{
			Namespace:          o.Namespace,
			ServiceAccountName: o.OCIServiceAccount,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get keychain: %w", err)
		}
		return oci.NewResolver(bundle, kc), nil
	}
	return nil, errors.Errorf("TODO")
}
