package code_registry

import (
	"fmt"
	"testing"
)

func TestRegistry(t *testing.T) {
	//gitlab.NewClient("GXEMyqUw1sJ3Ny3kGqFc", gitlab.WithBaseURL("http://121.5.102.204:8929/api/v4"))

	//client, _ := gitee.NewClient("fba6d0805b097e9ee05ae120c9d8ec9b", "")
	//client.RepoList()

	//client, _ := github.NewClient("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")
	//client.RepoList()

	client, err := NewCodeRegistry("gitee", "", "fba6d0805b097e9ee05ae120c9d8ec9b")

	//client, err := NewCodeRegistry("github", "", "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")

	//client, err := NewCodeRegistry("gitlab", "http://121.5.102.204:8929/api/v4", "GXEMyqUw1sJ3Ny3kGqFc")
	if err != nil {
		panic(err)
	}
	//project, err := client.GetRepo("hello", "hellotest1") //gitlab
	//project, err := client.GetRepo("limeschool", "gin") //github
	project, err := client.GetRepo("ptl-f", "ps-go") //github

	if err != nil {
		panic(err)
	}
	tags, err := client.GetRepoBranches(project)
	if err != nil {
		panic(err)
	}
	fmt.Println(tags)
}
