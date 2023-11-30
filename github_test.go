package packagebuku

import (
	"context"
	"fmt"
	"testing"
)

var PersonalToken = "personal"
var Reponame = "namerepo"
var OwnerName = "nameowner"

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
