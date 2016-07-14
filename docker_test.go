package ampsvc

import (
	"testing"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestDocker(t *testing.T) {
	dCli, _ := client.NewEnvClient()
	opts := types.ImageListOptions{}
	images, _ := dCli.ImageList(context.Background(), opts)
	t.Log(images)
	assert.Equal(t, "this", "this", "Temporary")
}
