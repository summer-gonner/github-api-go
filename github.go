package github

import (
	"github.com/summer-gonner/github/lang"
)

type Github struct {
	AccessToken string
	Language    string
}

var i18n *lang.Translation
