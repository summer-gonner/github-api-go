package client

import (
	"context"
	"github.com/google/go-github/v45/github"
)

type GithubRepository struct {
	Name        string
	Description string
	Private     bool
}

// GetRepositoryList 获取仓库列表
func (g Github) GetRepositoryList(owner string) ([]*github.Repository, error) {
	repositories, _, err := g.Client.Repositories.List(context.Background(), owner, nil)
	if err != nil {
		return nil, err
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
