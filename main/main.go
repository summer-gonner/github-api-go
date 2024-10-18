package main

import (
	"github.com/summer-gonner/github"
	"log"
)

var (
	owner      = "summer-gonner"
	repoName   = "k8s"
	branchName = "main"
	token      = "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"
)

func main() {
	g := &github.Github{AccessToken: token}
	newClient, err := github.NewClient(g)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf("11%s\n", newClient.Repositories)
}
