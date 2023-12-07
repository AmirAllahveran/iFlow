package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/smf8/http-monitor/model"
)

func (d *Client) ContainersList() (containers []*model.Container) {
	rawContainers, err := d.cli.ContainerList(context.TODO(), types.ContainerListOptions{All: true})
	fmt.Println(err)
	for _, rawCont := range rawContainers {
		containers = append(containers, &model.Container{
			ID:     rawCont.ID,
			Name:   rawCont.Names[0],
			Image:  rawCont.Image,
			Status: rawCont.Status,
			State:  model.ContainerState(rawCont.State),
		})
	}
	return
}

func (d *Client) GetContainer(containerID string) (*model.Container, error) {
	containerInspect, err := d.cli.ContainerInspect(context.TODO(), containerID)
	if err != nil {
		return nil, err
	}
	return &model.Container{
		ID:     containerInspect.ID,
		Name:   containerInspect.Name,
		Image:  containerInspect.Image,
		Status: containerInspect.State.Status,
		State:  model.ContainerState(containerInspect.State.Status),
	}, nil
}

func (d *Client) ContainerStats(containerID string) (io.ReadCloser, error) {
	stats, err := d.cli.ContainerStats(context.TODO(), containerID, true)
	return stats.Body, err
}

func (d *Client) ContainerCreate(config container.Config, networkID string, name string) (container.CreateResponse, error) {
	// Create container
	// Network configuration
	networkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			networkID: {},
		},
	}
	return d.cli.ContainerCreate(context.TODO(), &config, nil, networkingConfig, nil, "")

}

func (d *Client) ContainerStart(containerID string) error {
	return d.cli.ContainerStart(context.TODO(), containerID, types.ContainerStartOptions{})
}
