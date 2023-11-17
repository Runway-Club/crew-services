package testtools

import (
	"fmt"
	"time"
)

func NewElasticsearch() (dockerTest *DockerTest, dsn string) {
	dockerTest = NewDockerTest()
	dockerTest.CreateResource("elasticsearch", "8.11.0", []string{"xpack.security.enabled=false", "discovery.type=single-node"}, []string{}, 5*time.Second)
	port := dockerTest.Resource.GetPort("9200/tcp")
	dsn = fmt.Sprintf("http://localhost:%s", port)
	return dockerTest, dsn
}
