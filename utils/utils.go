package utils

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/devetek/go-git/types"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// GetURL should be used to ensure the right environment provide, return git repository url
func GetURL() string {
	URL := os.Getenv("GITHUB_URL")

	if URL == "" {
		Warning("GITHUB_URL not defined")
		os.Exit(1)
	}

	return URL
}

// GetDirectory should be used to ensure the right environment provide, return git target clone directory
func GetDirectory() string {
	dir, err := os.Getwd()

	if err != nil {
		Warning(err.Error())
		os.Exit(1)
	}

	directory := os.Getenv("GITHUB_DIRECTORY")

	if directory == "" {
		Warning("GITHUB_DIRECTORY not defined")
		os.Exit(1)
	}

	return path.Join(dir, directory)
}

// GetBranch should be used to ensure the right environment provide, return git target branch
func GetBranch() string {
	branch := os.Getenv("GITHUB_BRANCH")

	if branch == "" {
		Warning("GITHUB_BRANCH not defined")
		os.Exit(1)
	}

	return branch
}

// GetUsername should be used to ensure the right environment provide, return git username
func GetUsername() string {
	username := os.Getenv("GITHUB_USERNAME")

	if username == "" {
		Warning("GITHUB_USERNAME not defined")
		os.Exit(1)
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

// GetGitCloneOpt should be used to ensure the right environment provide, return git option
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

// GetGitPullOpt should be used to ensure the right environment provide, return git option
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

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(arg ...string) {
	if len(arg) == 0 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
