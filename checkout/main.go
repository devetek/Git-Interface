package main

import (
	"github.com/devetek/go-git/utils"

	"github.com/go-git/go-git/v5"
)

func main() {
	// opts := utils.GetGitPullOpt()
	directory := utils.GetDirectory()
	branch := utils.GetBranch()
	checkoutOpts, err := utils.GetGitCheckoutOpt("", branch)

	// retrieving error, check in utils and then print error if error not nil
	utils.CheckIfError(err)

	utils.Info("Go to directory" + directory)

	r, err := git.PlainOpen(utils.GetDirectory())

	// retrieving error, check in utils and then print error if error not nil
	utils.CheckIfError(err)

	w, err := r.Worktree()

	// retrieving error, check in utils and then print error if error not nil
	utils.CheckIfError(err)

	// Checkout the latest changes from the origin remote and merge into the current branch
	utils.Info("git checkout %s", branch)
	err = w.Checkout(checkoutOpts)

	utils.CheckIfError(err)

	ref, err := r.Head()

	// retrieving error, check in utils and then print error if error not nil
	utils.CheckIfError(err)

	utils.Info("Latest hash from current branch %s", ref)
}
