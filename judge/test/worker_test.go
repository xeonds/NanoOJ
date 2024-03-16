package test

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"xyz.xeonds/nano-oj/worker"
)

func TestConnToLocalDocker(t *testing.T) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Error("Failed to establish connection")
	}
	defer client.Close()
	t.Log(client.ImageList(context.Background(), types.ImageListOptions{}))
}

func TestConnToHttpDocker(t *testing.T) {
	client, err := worker.NewJudgeServer("unix:///var/run/docker.sock")
	if err != nil {
		t.Error("Failed to establish connection")
	}
	t.Log(client.ListImages())
}
