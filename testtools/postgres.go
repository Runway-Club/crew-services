package testtools

import (
	"fmt"
	"time"
)

func NewPostgresTest() (dockerTest *DockerTest, dsn string) {
	dockerTest = NewDockerTest()
	dockerTest.CreateResource("postgres", "13.1-alpine", []string{"POSTGRES_USER=test", "POSTGRES_PASSWORD=test"}, []string{}, 5*time.Second)
	port := dockerTest.Resource.GetPort("5432/tcp")
	dsn = fmt.Sprintf("host=localhost profile=test password=test dbname=test port=%s sslmode=disable TimeZone=Asia/Saigon", port)
	return dockerTest, dsn
}
