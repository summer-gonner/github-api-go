package github

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	githubv45 "github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
	"log"
)

type Client struct {
	Users        *UsersService
	Repositories *RepositoriesService
}

type Service struct {
	Users        *UsersService
	Repositories *RepositoriesService
}

// NewClient 初始化一个客户端可以进行配置
func NewClient(github *Github) (*Service, error) {
	if !govalidator.IsNull(github.Language) {
		i18n.Language = github.Language
	}
	if govalidator.IsNull(github.AccessToken) {
		return nil, fmt.Errorf(i18n.Translate("AccessToken不能为空"))
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: github.AccessToken,
		})
	tc := oauth2.NewClient(context.Background(), ts)
	client := githubv45.NewClient(tc)
	log.Printf("c%s\n", client)
	return &Service{
		Client: client,
	}, nil
}
