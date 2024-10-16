package client

import (
	"github.com/google/go-github/v45/github"
	"github.com/summer-gonner/github-api-go/lang"
)

type Github struct {
	AccessToken string
	Owner       string
	Client      *github.Client
	I18n        *lang.Translation
	User
}
