package nacos

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/go-kratos/kratos/v2/registry"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

// Registry is nacos registry.
type Registry struct {
	opts options
	cli  naming_client.INamingClient
}

// New new a nacos registry.
func New(cli naming_client.INamingClient, opts ...Option) (r *Registry) {
	op := options{
		prefix:  "/microservices",
		cluster: "DEFAULT",
		group:   constant.DEFAULT_GROUP,
		weight:  100,
		kind:    "grpc",
	}
	for _, option := range opts {
		option(&op)
	}
	return &Registry{
		opts: op,
		cli:  cli,
	}
}

// Register the registration.
func (r *Registry) Register(_ context.Context, si *registry.ServiceInstance) error {
	if si.Name == "" {
		return fmt.Errorf("kratos/nacos: serviceInstance.name can not be empty")
	}
	for _, endpoint := range si.Endpoints {
		u, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		host, port, err := net.SplitHostPort(u.Host)
		if err != nil {
			return err
		}
		p, err := strconv.Atoi(port)
		if err != nil {
			return err
		}
		var rmd map[string]string
		if si.Metadata == nil {
			rmd = map[string]string{
				"kind":    u.Scheme,
				"version": si.Version,
			}
		} else {
			rmd = make(map[string]string, len(si.Metadata)+2)
			for k, v := range si.Metadata {
				rmd[k] = v
			}
			rmd["kind"] = u.Scheme
			rmd["version"] = si.Version
		}
		_, e := r.cli.RegisterInstance(vo.RegisterInstanceParam{
			Ip:          host,
			Port:        uint64(p),
			ServiceName: si.Name + "." + u.Scheme,
			Weight:      r.opts.weight,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			Metadata:    rmd,
			ClusterName: r.opts.cluster,
			GroupName:   r.opts.group,
		})
		if e != nil {
			return fmt.Errorf("RegisterInstance err %v,%v", e, endpoint)
		}
	}
	return nil
}

// Deregister the registration.
func (r *Registry) Deregister(_ context.Context, service *registry.ServiceInstance) error {
	for _, endpoint := range service.Endpoints {
		u, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		host, port, err := net.SplitHostPort(u.Host)
		if err != nil {
			return err
		}
		p, err := strconv.Atoi(port)
		if err != nil {
			return err
		}
		if _, err = r.cli.DeregisterInstance(vo.DeregisterInstanceParam{
			Ip:          host,
			Port:        uint64(p),
			ServiceName: service.Name + "." + u.Scheme,
			GroupName:   r.opts.group,
			Cluster:     r.opts.cluster,
			Ephemeral:   true,
		}); err != nil {
			return err
		}
	}
	return nil
}

// Watch creates a watcher according to the service name.
func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	return newWatcher(ctx, r.cli, serviceName, r.opts.group, r.opts.kind, []string{r.opts.cluster})
}

// GetService return the service instances in memory according to the service name.
func (r *Registry) GetService(_ context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	res, err := r.cli.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   r.opts.group,
		HealthyOnly: true,
	})
	if err != nil {
		return nil, err
	}
	items := make([]*registry.ServiceInstance, 0, len(res))
	for _, in := range res {
		kind := r.opts.kind
		if k, ok := in.Metadata["kind"]; ok {
			kind = k
		}
		items = append(items, &registry.ServiceInstance{
			ID:        in.InstanceId,
			Name:      in.ServiceName,
			Version:   in.Metadata["version"],
			Metadata:  in.Metadata,
			Endpoints: []string{fmt.Sprintf("%s://%s:%d", kind, in.Ip, in.Port)},
		})
	}
	return items, nil
}
