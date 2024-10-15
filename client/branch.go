package client

import (
	"context"
	"github.com/google/go-github/v45/github"
	"log"
)

func (g Github) GetBranchList(owner string, repoName string) ([]*github.Branch, error) {
	branches, response, err := g.Client.Repositories.ListBranches(context.Background(), owner, repoName, nil)
	if err != nil {
		log.Printf("GetBranchList  Response:%+v\n", response)
		return nil, err
	}
	return branches, nil
}
