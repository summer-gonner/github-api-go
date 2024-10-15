package main

import (
	"github.com/summer-gonner/github-api-go/client"
	"github.com/summer-gonner/github-api-go/lang"
	"log"
)

var (
	owner      = "牛逼plus"
	repoName   = "k8s"
	branchName = "main"
)

func main() {

	githubInfo := client.Github{
		AccessToken: "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",
		Lang:        lang.English,
	}

	newClient, err := githubInfo.NewClient()
	if err != nil {
		log.Println("NewClient", err)
		return
	}
	user, err := newClient.GetUser(owner)
	if err != nil {
		log.Println("获取用户失败：", err)
		return
	}
	log.Printf("获取用户：%s", user)

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
