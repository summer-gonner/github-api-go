package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/summer-gonner/github/http"
)

type User struct {
	Name string
}

type UsersService Service

func (u UsersService) GetUser(name string) (*github.User, error) {
	user, response, err := u.client.Users.Get(context.Background(), name)
	if response.Status != http.Success || err != nil {
		_, err := http.IoToErrorResponse(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(i18n.Translate("用户不存在!"))
	}
	return user, nil
}
