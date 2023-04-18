package worker

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func InitJudgerImage() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation()) // initiate docker connection
	if err != nil {
		log.Fatal("failed to connect to docker daemon", err)
	}
	defer cli.Close()

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{}) // find whether the judge image is exist
	if err != nil {
		log.Fatal("failed to list docker images", err)
	}
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == "nano-oj/judge" {
				return
			}
		}
	}

	_, err = cli.ImageBuild(context.Background(), nil, types.ImageBuildOptions{ // if the judge image isn't exist, build it from the Dockerfile
		Dockerfile: "judge.Dockerfile",
		Tags:       []string{"nano-oj/judge"},
	})
	if err != nil {
		log.Fatal("failed to build judge image", err)
	}
}
