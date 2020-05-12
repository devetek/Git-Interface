package utils

import (
	"os"
	"path"
	"testing"

	"github.com/devetek/go-git/types"
)

func TestGetURL(t *testing.T) {
	var testURL = "https://github.com/devetek/Simple-Kaniko.git"
	os.Setenv("GITHUB_URL", testURL)

	url := GetURL()

	if url != testURL {
		t.Errorf("got %q want %q", url, testURL)
	}
}

func TestGetDirectory(t *testing.T) {
	dir, _ := os.Getwd()
	var testDir = "git-test"
	os.Setenv("GITHUB_DIRECTORY", testDir)

	var targetDir = path.Join(dir, testDir)

	directory := GetDirectory()

	if directory != targetDir {
		t.Errorf("got %q want %q", directory, targetDir)
	}
}

func TestGetUsername(t *testing.T) {
	var testUsername = "prakasa1904"
	os.Setenv("GITHUB_USERNAME", testUsername)

	username := GetUsername()

	if username != testUsername {
		t.Errorf("got %q want %q", username, testUsername)
	}
}

func TestGetToken(t *testing.T) {
	var testToken = ""
	os.Setenv("GITHUB_TOKEN", testToken)

	token := GetToken()

	if token != testToken {
		t.Errorf("got %q want %q", token, testToken)
	}
}

func TestCheckArgs(t *testing.T) {
	CheckArgs("<url>", "<directory>", "<username>", "<token>")
}

func TestInfo(t *testing.T) {
	var env types.GithubEnv

	env = types.GithubEnv{
		URL:       "https://github.com/devetek/Simple-Kaniko.git",
		Directory: "./git-kaniko",
		Username:  "prakasa1904",
	}

	Info("git clone %s %s --recursive", env.URL, env.Directory)
}

func TestWarning(t *testing.T) {
	Warning("It just test from warning function")
}

func TestCheckIfError(t *testing.T) {
	var notError error = nil

	CheckIfError(notError)
}
