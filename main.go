package main

import (
	"github.com/summer-gonner/github-api-go/client"
	"github.com/summer-gonner/github-api-go/lang"
	"log"
)

var (
	owner      = "summer-gonner"
	repoName   = "k8s"
	branchName = "main"
)

func main() {

	githubInfo := client.Github{
		AccessToken: "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",
		I18n: &lang.Translation{
			Language: lang.English,
		},
	}

	newClient, err := githubInfo.NewClient()
	if err != nil {
		log.Println("NewClient", err)
		return
	}
	//repositoryList, err := newClient.GetRepositoryList(owner)
	//if err != nil {
	//	log.Println("err：", err)
	//	return
	//}
	//log.Printf("获取仓库：%s", repositoryList)
	newClient.Cre

	//ref, err := newClient.GetGitRef(owner, repoName, branchName)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Printf("分支信息%+v\n", ref)
	//
	//gitRef, err := newClient.CreateGitRef(owner, repoName, branchName, "uat")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Printf("创建git分支%+v\n", gitRef)

}
