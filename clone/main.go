package main

import (
	"github.com/devetek/go-git/utils"

	"github.com/go-git/go-git/v5"
)

func main() {
	opts := utils.GetGitCloneOpt()
	directory := utils.GetDirectory()
	URL := utils.GetURL()

	utils.Info("git clone %s %s --recursive", URL, directory)

	_, err := git.PlainClone(utils.GetDirectory(), false, &opts)

	// retrieving error, check in utils and then print error if error not nil
	utils.CheckIfError(err)
}
