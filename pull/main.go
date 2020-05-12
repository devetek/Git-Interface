package main

import (
	"github.com/devetek/go-git/utils"

	"github.com/go-git/go-git/v5"
)

func main() {
	opts := utils.GetGitPullOpt()
	directory := utils.GetDirectory()

	utils.Info("Go to directory", directory)

	r, err := git.PlainOpen(utils.GetDirectory())

	// retrieving error, check in utils and then print error if error not nil
	utils.CheckIfError(err)

	// Get the working directory for the repository
	w, err := r.Worktree()

	// retrieving error, check in utils and then print error if error not nil
	utils.CheckIfError(err)

	// Pull the latest changes from the origin remote and merge into the current branch
	utils.Info("git pull origin")

	err = w.Pull(&opts)

	// retrieving error, check in utils and then print error if error not nil
	utils.CheckIfError(err)
}
