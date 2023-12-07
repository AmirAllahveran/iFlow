package docker

import (
	"io"

	"github.com/arshamalh/dockeroller/log"
	"github.com/arshamalh/dockeroller/models"
	"github.com/moby/moby/client"
)

type Docker interface {
	GetContainer(containerID string) (*models.Container, error)
	ContainersList() []*models.Container
	ContainerLogs(containerID string) (io.ReadCloser, error)
	ContainerStats(containerID string) (io.ReadCloser, error)
	ContainerStart(containerID string) error
	ImagePull(imageName string) error
}

func NewDocker() *Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Gl.Error(err.Error())
	}
	return &Client{
		cli: cli,
	}
}

type Client struct {
	cli *client.Client
}
