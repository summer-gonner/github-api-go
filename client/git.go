package client

import (
	"context"
	"github.com/google/go-github/v45/github"
	"github.com/summer-gonner/github-api-go/http"
	"log"
)

// GetGitRef  获取基础分支的详细信息
func (g Github) GetGitRef(owner string, repoName string, branchName string) (*github.Reference, error) {
	ref, _, err := g.Client.Git.GetRef(context.Background(), owner, repoName, "refs/heads/"+branchName)
	if err != nil {
		//log.Printf("GetGitRef Response:%+v\n", response)
		return nil, err
	}
	return ref, nil
}

// CreateGitRef 创建新的分支引用
func (g Github) CreateGitRef(owner string, repoName string, baseBranch string, newRef string) (*github.Reference, error) {
	//获取基础分支的详细信息
	ref, _, err := g.Client.Git.GetRef(context.Background(), owner, repoName, "refs/heads/"+baseBranch)
	if err != nil {
		//log.Printf("GetGitRef Response:%+v\n", response)
		return nil, err
	}
	newReference := &github.Reference{
		Ref:    github.String("refs/heads/" + newRef),
		Object: &github.GitObject{SHA: ref.Object.SHA},
	}
	//创建新的分支引用
	createRef, response, err := g.Client.Git.CreateRef(context.Background(), owner, repoName, newReference)
	log.Printf("CreateGitRef Response:%+v\n", response.Response.Status)
	if err != nil {

		errorResponse, err := http.IoToErrorResponse(response.Response.Body)
		log.Printf("errorResponse:%+v\n", errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return createRef, nil
}
