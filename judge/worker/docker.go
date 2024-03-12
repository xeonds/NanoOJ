package worker

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/database/model"
	"xyz.xeonds/nano-oj/utils"
)

// init judger pool when the server starts
var JudgerPool = make(map[string]JudgeServer)

type JudgeServer struct {
	db    *gorm.DB
	cli   *client.Client
	ctx   context.Context
	Works chan task
}

type IJudgeServer interface {
	ListImages() ([]types.ImageSummary, error)
	BuildImage() error
	RunTask(t task) error
}

func InitJudgerPool() {
	judgers := utils.GetJudgers()
	for _, judger := range judgers {
		judgerServer, err := NewJudgeServer(judger)
		if err != nil {
			log.Println("failed to connect to judger ", judger, ": ", err)
			continue
		}
		JudgerPool[judger] = judgerServer
	}
}

func (c JudgeServer) AddTask(t task) {
	c.Works <- t
}

// TODO: should be verified
// starts all judgers in the pool using goroutine, and
// continuously listen to the task channel
func RunJudgerPool() {
	for _, judger := range JudgerPool {
		go func(judger JudgeServer) {
			judger.Process()
		}(judger)
	}
}

// TODO: test the function. it's just a demo now
func NewJudgeServer(addr string) (JudgeServer, error) {
	cli, err := client.NewClient(addr, "", nil, nil) // initiate docker connection
	if err != nil {
		log.Println("failed to connect to docker daemon ", addr, ": ", err)
		return JudgeServer{}, err
	}
	ctx := context.Background()
	return JudgeServer{cli: cli, ctx: ctx}, nil
}

// get a judger from the judger pool with the least load
func GetAvailableJudger() JudgeServer {
	var minLoad = 10000
	var minLoadJudger JudgeServer
	for _, judger := range JudgerPool {
		load := len((JudgeServer)(judger).Works)
		if load < minLoad {
			minLoad = load
			minLoadJudger = judger
		}
	}
	return minLoadJudger
}

// list all images at the target server
func (c JudgeServer) ListImages() ([]types.ImageSummary, error) {
	images, err := c.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.Fatal("failed to list docker images", err)
	}
	return images, err
}

// TODO: this function needs improvement: needs to create the image from the given Dockerfile(embedded in the binary file)
// builds the judge image from the Dockerfile
func (c JudgeServer) BuildImage() error {
	images, err := c.cli.ImageList(context.Background(), types.ImageListOptions{}) // find whether the judge image is exist
	if err != nil {
		log.Fatal("failed to list docker images", err)
	}
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == "nano-oj/judge" {
				return nil
			}
		}
	}

	_, err = c.cli.ImageBuild(context.Background(), nil, types.ImageBuildOptions{ // if the judge image isn't exist, build it from the Dockerfile
		Dockerfile: "judge.Dockerfile",
		Tags:       []string{"nano-oj/judge"},
	})
	if err != nil {
		log.Println("failed to build judge image", err)
		return err
	}
	return nil
}

// TODO: needs to be tested, and improve: the image used to test should be chosen by language type
// run a task in a container
func (c JudgeServer) RunTask(t task) (result, error) {
	// create a container from the judge image, give it the workdir and source file as volume and run it
	resp, err := c.cli.ContainerCreate(c.ctx, &container.Config{
		Image:           "nano-oj/judge",
		Env:             []string{fmt.Sprintf("LANG=%s", t.Lang), fmt.Sprintf("TIME_LIMIT=%d", t.TimeLimit)},
		NetworkDisabled: true,
	}, &container.HostConfig{
		Binds: []string{fmt.Sprintf("%s:/code", t.Workdir)},
	}, nil, nil, "")
	defer os.RemoveAll(t.Workdir)
	if err != nil {
		return result{model.CompilationError, "Failed to create container"}, err
	}
	err = c.cli.ContainerStart(c.ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return result{model.CompilationError, "Failed to start container"}, err
	}
	// wait for the container to finish
	statusCh, errCh := c.cli.ContainerWait(c.ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return result{model.CompilationError, "Failed to wait container"}, err
		}
	case <-statusCh:
	}
	// read the result from the container
	out, err := c.cli.ContainerLogs(c.ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return result{model.CompilationError, "Failed to read container logs"}, err
	}
	// delete the container
	err = c.cli.ContainerRemove(c.ctx, resp.ID, types.ContainerRemoveOptions{})
	if err != nil {
		return result{model.CompilationError, "Failed to remove container"}, err
	}
	// parse the result
	stat, info, err := utils.ParseResult(out)
	if err != nil {
		return result{model.CompilationError, "Failed to parse result"}, err
	}
	return result{stat, info}, nil
}

// process all tasks in the queue, using goroutine
func (c JudgeServer) Process() {
	for {
		t := <-c.Works
		// get a task from the queue
		res, err := c.RunTask(t)
		if err != nil {
			log.Println("failed to run task", err)
		}
		// update the result to the database
		CommitStatus(c.db, t.Submission, res.Status, res.Info)
	}
}
