package client

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/summer-gonner/github-api-go/http"
	"log"
)

// GetUser 获取某个用户信息
func (g Github) GetUser(owner string) (*github.User, error) {
	log.Printf("语言%s\n", g.Lang)
	user, response, err := g.Client.Users.Get(context.Background(), owner)
	if response.Status != http.Success || err != nil {
		_, err := http.IoToErrorResponse(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(g.I18n("用户不存在!"))
	}
	return user, nil
}
