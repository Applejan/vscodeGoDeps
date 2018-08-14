package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

//g save the url and the target dir
type g struct {
	url   string
	toDir string
}

func downTOOLS() {
	tools := g{
		url:   "https://github.com/golang/tools",
		toDir: "$GOPATH/src/golang.org/x/tools",
	}
	k := filepath.SplitList(tools.toDir)
	saveDir := os.Getenv("GOPATH") + filepath.Join(k...)
	cmd := exec.Command("git", "pull", tools.url, k...)
}
