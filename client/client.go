package client

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

// NewClient 初始化一个客户端可以进行配置
func (g Github) NewClient() (*Github, error) {
	if govalidator.IsNull(g.AccessToken) {
		return nil, fmt.Errorf(g.I18n.Translate("AccessToken不能为空"))
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: g.AccessToken,
		})
	tc := oauth2.NewClient(context.Background(), ts)
	return &Github{
		Client: github.NewClient(tc),
		I18n:   g.I18n,
	}, nil
}
