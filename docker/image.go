package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"io"
)

func (d *Client) ImagePull(imageName string) error {
	// Pull the image
	out, err := d.cli.ImagePull(context.TODO(), imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer func(out io.ReadCloser) {
		err := out.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(out)

	return nil
}
