package packagebuku

import (
	"context"
	"fmt"
	"testing"
)

var PersonalToken = "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"
var Reponame = "testuploadgithub"
var OwnerName = "HRMonitorr"

func TestUploadFileToRepository(t *testing.T) {
	value := PushRepositories{
		Context:       context.Background(),
		PersonalToken: PersonalToken,
		Reponame:      Reponame,
		OwnerName:     OwnerName,
		Path:          "namafile.txt",
		Username:      "username",
		Email:         "email@gmail.com",
		Message:       "Percobaan test push dari golang",
		Branch:        "master",
	}
	push, err := UploadFileToRepository(value)
	fmt.Printf("err %+v\n", err)
	fmt.Printf("err %+v\n", push.Message)
}

func TestGetListRepositories(t *testing.T) {
	list, err := ListRepositoriesOrg(context.Background(),
		PersonalToken,
		OwnerName,
	)
	fmt.Printf("%+v\n", list)
	fmt.Printf("%+v\n", err)
}

func TestGetCommitAll(t *testing.T) {
	url, err := ListCommitALL(context.Background(),
		PersonalToken,
		Reponame,
		OwnerName)
	fmt.Printf("%+v\n", url)
	fmt.Printf("%+v\n", err)
	//fmt.Printf("%+v\n", comms)
}

func TestGetListRepositoriesDetail(t *testing.T) {
	Det, err := ListRepositoriesOnlydDetail(context.Background(),
		PersonalToken,
		OwnerName,
	)
	fmt.Printf("%+v\n", *Det[0].Name)
	fmt.Printf("%+v\n", err)
}
