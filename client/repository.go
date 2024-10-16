package client

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/summer-gonner/github-api-go/http"
)

type GithubRepository struct {
	Name        string
	Description string
	Private     bool
}

// GetRepositoryList 获取仓库列表
func (g Github) GetRepositoryList(owner string) ([]*github.Repository, error) {
	repositories, res, err := g.Client.Repositories.List(context.Background(), owner, nil)
	if res.Status != http.Success || err != nil {
		_, err := http.IoToErrorResponse(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(g.I18n.Translate("仓库列表为空!"))
	}
	return repositories, nil
}

// CreateRepository 创建仓库
func (g Github) CreateRepository(gr GithubRepository) (*github.Repository, error) {
	repo := &github.Repository{
		Name:        github.String(gr.Name),
		Description: github.String(gr.Description),
		Private:     github.Bool(gr.Private),
	}
	createRepo, _, err := g.Client.Repositories.Create(context.Background(), "", repo)
	if err != nil {
		return nil, err
	}
	return createRepo, nil
}
func (g Github) DeleteRepository(owner string, repoName string) error {
	_, err := g.Client.Repositories.Delete(context.Background(), owner, repoName)
	if err != nil {
		return err
	}
	return nil
}

func (g Github) EditRepository(owner string, oldRepoName string, gr GithubRepository) (*github.Repository, error) {
	repo := &github.Repository{
		Name: github.String(gr.Name),
	}
	editRepository, _, err := g.Client.Repositories.Edit(context.Background(), owner, oldRepoName, repo)
	if err != nil {
		return nil, err
	}
	return editRepository, nil
}
