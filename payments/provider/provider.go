package provider

import (
	"errors"

	"github.com/micro/go-micro/v3/registry"
	mregistry "github.com/micro/micro/v3/service/registry"

	pb "github.com/m3o/services/payments/provider/proto"
)

// Provider is an alias type so clients don't need to also import the pb
type Provider = pb.ProviderService

// ServicePrefix is the prefix appended to a provider name to get
// the service type
const ServicePrefix = "go.micro.service.payment."

var (
	// ErrNotFound is returned when a provider is not found in the registry
	ErrNotFound = errors.New("Provider not found")
)

// NewProvider returns an initialized client with the name provided,
// e.g. "stripe" will return a payment provider with the service name
// "go.micro.service.payments.stripe"
func NewProvider(name string) (pb.ProviderService, error) {
	// Construct the service name
	srvName := ServicePrefix + name

	// Check the service exists in the registry (ensuring we fail fast if not)
	srvs, err := mregistry.DefaultRegistry.GetService(srvName)
	if len(srvs) == 0 || err == registry.ErrNotFound {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// Return an initialized provider service
	srv := pb.NewProviderService(srvName)
	return srv, nil
}
