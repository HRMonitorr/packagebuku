package packagebuku

import (
	"bufio"
	"fmt"
	"github.com/google/go-github/v56/github"
	"io"
	"log"
	"os"
)

func MakeClient(personalToken string) *github.Client {
	client := github.NewClient(nil).WithAuthToken(personalToken)
	return client
}

func UploadFileToRepository(val PushRepositories) (response *github.RepositoryContentResponse, err error) {
	file, err := os.Open(val.Path)
	if err != nil {
		log.Fatalf("Gagal open file %s", err.Error())
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String(val.Message),
		Content:   bs,
		Branch:    github.String(val.Branch),
		Committer: &github.CommitAuthor{Name: github.String(val.Username), Email: github.String(val.Email)},
	}
	response, _, err = MakeClient(val.PersonalToken).Repositories.CreateFile(val.Context, val.OwnerName, val.Reponame, val.Path, opts)
	if err != nil {
		fmt.Printf("%+v", err.Error())
		return
	}
	return
}
