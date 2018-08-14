package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//g save the url and the target dir
type body struct {
	url   string
	toDir string
}

func (s *body) downTOOLS() {
	tmp := strings.Split(s.toDir, "/")
	tmp = tmp[1:]
	// log.Println(tmQp)
	saveDir := os.Getenv("GOPATH") + "\\" + filepath.Join(tmp...)
	args := []string{"pull", saveDir}
	// log.Println(args)
	cmd := exec.Command("git", args...)
	cmd.Run()
}
func main() {

	tools := body{
		url:   "https://github.com/golang/tools",
		toDir: "$GOPATH/src/golang.org/x/tools",
	}
	tools.downTOOLS()
}
