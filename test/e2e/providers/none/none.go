package none

import (
	"fmt"

	config "github.com/openshift/api/config/v1"
	mapi "github.com/openshift/api/machine/v1beta1"
	core "k8s.io/api/core/v1"
	client "k8s.io/client-go/kubernetes"

	"github.com/openshift/windows-machine-config-operator/test/e2e/clusterinfo"
	"github.com/openshift/windows-machine-config-operator/test/e2e/windows"
)

// Provider is a provider struct for testing platform=none
type Provider struct {
	oc *clusterinfo.OpenShift
}

// New returns a Provider implementation for platform=none
func New(clientset *clusterinfo.OpenShift) (*Provider, error) {
	return &Provider{
		oc: clientset,
	}, nil
}

// GenerateMachineSet is not supported for platform=none and throws an exception
func (p *Provider) GenerateMachineSet(_ bool, replicas int32, version windows.ServerVersion) (*mapi.MachineSet, error) {
	return nil, fmt.Errorf("MachineSet generation not supported for platform=none")
}

// GetType returns the platform type for platform=none
func (p *Provider) GetType() config.PlatformType {
	return config.NonePlatformType
}

func (p *Provider) StorageSupport() bool {
	return false
}

func (p *Provider) CreatePVC(_ client.Interface, _ string) (*core.PersistentVolumeClaim, error) {
	return nil, fmt.Errorf("storage not supported on platform none")
}
