package main

import (
	"os"

	"github.com/devetek/go-git/utils"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {
	utils.CheckArgs("<url>", "<directory>", "<username>", "<token>")
	url, directory, username, token := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	utils.Info("git clone %s %s --recursive", url, directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
		Depth:             1,
		Auth: &http.BasicAuth{
			Username: username, // yes, this can be anything except an empty string
			Password: token,
		},
	})

	// retrieving error, check in utils and then print error if error not nil
	utils.CheckIfError(err)
}
