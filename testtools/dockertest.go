package testtools

import (
	"time"

	"github.com/ory/dockertest/v3"
)

type DockerTest struct {
	Pool     *dockertest.Pool
	Resource *dockertest.Resource
}

func NewDockerTest() *DockerTest {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}
	return &DockerTest{pool, nil}
}

func (d *DockerTest) CreateResource(image string, tag string, env []string, cmd []string, waitTime time.Duration) {
	resource, err := d.Pool.Run(image, tag, env)
	if err != nil {
		panic(err)
	}
	d.Resource = resource
	<-time.After(waitTime)
}

func (d *DockerTest) PurgeResource() {
	if err := d.Pool.Purge(d.Resource); err != nil {
		panic(err)
	}
}
