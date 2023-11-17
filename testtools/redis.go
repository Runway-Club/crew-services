package testtools

import (
	"fmt"
	"time"
)

func NewRedisTest(dbName int) (dockerTest *DockerTest, dsn string) {
	dockerTest = NewDockerTest()
	dockerTest.CreateResource("redis", "alpine3.18", []string{}, []string{}, 5*time.Second)
	port := dockerTest.Resource.GetPort("6379/tcp")
	dsn = fmt.Sprintf("localhost:%s", port)
	return dockerTest, dsn
}
