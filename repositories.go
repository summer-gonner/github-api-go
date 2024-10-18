package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/summer-gonner/github/http"
)

type Repository struct {
	Branch
}

type RepositoriesService Service

func (r RepositoriesService) List(username string) ([]*github.Repository, error) {
	repositories, response, err := r.client.Repositories.List(context.Background(), username, nil)
	if response.Status != http.Success || err != nil {
		_, err := http.IoToErrorResponse(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(i18n.Translate("仓库列表为空!"))
	}
	return repositories, nil
}
