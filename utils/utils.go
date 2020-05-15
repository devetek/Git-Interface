package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/devetek/go-git/types"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// GetURL should be used to ensure the right environment provide, return git repository url
func GetURL() string {
	URL := os.Getenv("GITHUB_URL")

	if URL == "" {
		Warning("GITHUB_URL not defined")
		os.Exit(0)
	}

	return URL
}

// GetDirectory should be used to ensure the right environment provide, return git target clone directory
func GetDirectory() string {
	dir, err := os.Getwd()

	if err != nil {
		Warning(err.Error())
		os.Exit(0)
	}

	directory := os.Getenv("GITHUB_DIRECTORY")

	if directory == "" {
		Warning("GITHUB_DIRECTORY not defined")
		os.Exit(0)
	}

	return path.Join(dir, directory)
}

// GetBranch should be used to ensure the right environment provide, return git target branch
func GetBranch() string {
	branch := os.Getenv("GITHUB_BRANCH")

	if branch == "" {
		Warning("GITHUB_BRANCH not defined")
		os.Exit(0)
	}

	return branch
}

// GetUsername should be used to ensure the right environment provide, return git username
func GetUsername() string {
	username := os.Getenv("GITHUB_USERNAME")

	if username == "" {
		Warning("GITHUB_USERNAME not defined")
		os.Exit(0)
	}

	return username
}

// GetToken should be used to ensure the right environment provide, return git token
func GetToken() string {
	token := os.Getenv("GITHUB_TOKEN")

	return token
}

// GetEnvironments should be used to ensure the right environment provide, return git token
func GetEnvironments() types.GithubEnv {
	var result types.GithubEnv

	result = types.GithubEnv{
		URL:       GetURL(),
		Directory: GetDirectory(),
		Username:  GetUsername(),
		Token:     GetToken(),
	}

	return result
}

// GetGitCloneOpt should be used to ensure the right environment provide, return git clone option
func GetGitCloneOpt() git.CloneOptions {
	var result git.CloneOptions

	url := GetURL()
	username := GetUsername()
	token := GetToken()

	if token != "" {
		result = git.CloneOptions{
			URL:               url,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			Progress:          os.Stdout,
			Auth: &http.BasicAuth{
				Username: username, // yes, this can be anything except an empty string
				Password: token,
			},
		}

		return result
	}

	result = git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	}

	return result
}

// GetGitFetchOpt should be used to ensure the right environment provide, return git fetch option
func GetGitFetchOpt() git.FetchOptions {
	var result git.FetchOptions

	username := GetUsername()
	token := GetToken()
	branch := GetBranch()

	if token != "" {
		result = git.FetchOptions{
			RefSpecs: []config.RefSpec{config.RefSpec("refs/heads/" + branch + ":refs/heads/" + branch)},
			Force:    true,
			Auth: &http.BasicAuth{
				Username: username, // yes, this can be anything except an empty string
				Password: token,
			},
			Progress: os.Stdout,
		}

		return result
	}

	result = git.FetchOptions{
		RefSpecs: []config.RefSpec{config.RefSpec("refs/heads/" + branch + ":refs/heads/" + branch)},
		Force:    true,
		Progress: os.Stdout,
	}

	return result
}

// GetGitPullOpt should be used to ensure the right environment provide, return git pull option
func GetGitPullOpt() git.PullOptions {
	var result git.PullOptions

	username := GetUsername()
	token := GetToken()

	if token != "" {
		result = git.PullOptions{
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			Progress:          os.Stdout,
			Auth: &http.BasicAuth{
				Username: username, // yes, this can be anything except an empty string
				Password: token,
			},
		}

		return result
	}

	result = git.PullOptions{
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	}

	return result
}

// GetGitCheckoutOpt should be used to ensure the right environment provide, return git checkout option
func GetGitCheckoutOpt(hash string, branch string) (*git.CheckoutOptions, error) {
	var result git.CheckoutOptions

	if hash == "" && branch == "" {
		return &result, errors.New("Invalid params")
	}

	if hash == "" && branch != "" {
		result = git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/heads/" + branch),
		}

	}

	if hash != "" && branch == "" {
		result = git.CheckoutOptions{
			Hash: plumbing.NewHash(hash),
		}

	}

	return &result, nil
}

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(arg ...string) {
	if len(arg) == 0 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(0)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(0)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
