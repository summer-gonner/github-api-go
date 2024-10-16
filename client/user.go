package client

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/summer-gonner/github-api-go/http"
)

// GetUser 获取某个用户信息
func (g Github) GetUser(owner string) (*github.User, error) {
	user, res, err := g.Client.Users.Get(context.Background(), owner)
	if res.Status != http.Success || err != nil {
		_, err := http.IoToErrorResponse(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(g.I18n.Translate("用户不存在!"))
	}
	return user, nil
}
