package utils

import (
	"fmt"
	"testing"

	"github.com/devetek/go-git/types"
)

func TestCheckArgs(t *testing.T) {
	var os types.Github
	os.Args = []string{"url", "./test", "prakasa1904", ""}
	fmt.Println(os.Args)

	CheckArgs("<url>", "<directory>", "<username>", "<token>")
}
