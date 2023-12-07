package docker

import (
	"context"
	"github.com/docker/docker/api/types"
)

func (d *Client) CreateNetwork(netConfig types.NetworkCreate, name string) (types.NetworkCreateResponse, error) {
	return d.cli.NetworkCreate(context.TODO(), name, netConfig)
}
