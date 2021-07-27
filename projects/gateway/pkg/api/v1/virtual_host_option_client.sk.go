// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type VirtualHostOptionWatcher interface {
	// watch namespace-scoped VirtualHostOptions
	Watch(namespace string, opts clients.WatchOpts) (<-chan VirtualHostOptionList, <-chan error, error)
}

type VirtualHostOptionClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*VirtualHostOption, error)
	Write(resource *VirtualHostOption, opts clients.WriteOpts) (*VirtualHostOption, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (VirtualHostOptionList, error)
	VirtualHostOptionWatcher
}

type virtualHostOptionClient struct {
	rc clients.ResourceClient
}

func NewVirtualHostOptionClient(ctx context.Context, rcFactory factory.ResourceClientFactory) (VirtualHostOptionClient, error) {
	return NewVirtualHostOptionClientWithToken(ctx, rcFactory, "")
}

func NewVirtualHostOptionClientWithToken(ctx context.Context, rcFactory factory.ResourceClientFactory, token string) (VirtualHostOptionClient, error) {
	rc, err := rcFactory.NewResourceClient(ctx, factory.NewResourceClientParams{
		ResourceType: &VirtualHostOption{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base VirtualHostOption resource client")
	}
	return NewVirtualHostOptionClientWithBase(rc), nil
}

func NewVirtualHostOptionClientWithBase(rc clients.ResourceClient) VirtualHostOptionClient {
	return &virtualHostOptionClient{
		rc: rc,
	}
}

func (client *virtualHostOptionClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *virtualHostOptionClient) Register() error {
	return client.rc.Register()
}

func (client *virtualHostOptionClient) Read(namespace, name string, opts clients.ReadOpts) (*VirtualHostOption, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*VirtualHostOption), nil
}

func (client *virtualHostOptionClient) Write(virtualHostOption *VirtualHostOption, opts clients.WriteOpts) (*VirtualHostOption, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(virtualHostOption, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*VirtualHostOption), nil
}

func (client *virtualHostOptionClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *virtualHostOptionClient) List(namespace string, opts clients.ListOpts) (VirtualHostOptionList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToVirtualHostOption(resourceList), nil
}

func (client *virtualHostOptionClient) Watch(namespace string, opts clients.WatchOpts) (<-chan VirtualHostOptionList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	virtualHostOptionsChan := make(chan VirtualHostOptionList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				select {
				case virtualHostOptionsChan <- convertToVirtualHostOption(resourceList):
				case <-opts.Ctx.Done():
					close(virtualHostOptionsChan)
					return
				}
			case <-opts.Ctx.Done():
				close(virtualHostOptionsChan)
				return
			}
		}
	}()
	return virtualHostOptionsChan, errs, nil
}

func convertToVirtualHostOption(resources resources.ResourceList) VirtualHostOptionList {
	var virtualHostOptionList VirtualHostOptionList
	for _, resource := range resources {
		virtualHostOptionList = append(virtualHostOptionList, resource.(*VirtualHostOption))
	}
	return virtualHostOptionList
}