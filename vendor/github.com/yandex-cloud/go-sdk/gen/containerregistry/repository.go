// Code generated by sdkgen. DO NOT EDIT.

//nolint
package containerregistry

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
)

//revive:disable

// RepositoryServiceClient is a containerregistry.RepositoryServiceClient with
// lazy GRPC connection initialization.
type RepositoryServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ containerregistry.RepositoryServiceClient = &RepositoryServiceClient{}

// List implements containerregistry.RepositoryServiceClient
func (c *RepositoryServiceClient) List(ctx context.Context, in *containerregistry.ListRepositoriesRequest, opts ...grpc.CallOption) (*containerregistry.ListRepositoriesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return containerregistry.NewRepositoryServiceClient(conn).List(ctx, in, opts...)
}
