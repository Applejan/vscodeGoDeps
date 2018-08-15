package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//g save the url and the target dir
type g struct {
	url   string
	toDir string
}

func download(content []g) {

	for _, v := range content {
		tmp := strings.Split(v.toDir, "/")
		tmp[0] = os.Getenv("GOPATH")
		targetDir := strings.Join(tmp, string(os.PathSeparator))
		os.Chdir(targetDir)
		cmd := exec.Command("git", "pull")
		ret, _ := cmd.Output()

		fmt.Printf("Current work in %v \n", targetDir)
		fmt.Println(string(ret))
	}

}
func main() {
	//Need to update files
	tools := g{
		url:   "https://github.com/golang/tools",
		toDir: "$GOPATH/src/golang.org/x/tools",
	}
	sys := g{
		url:   "https://github.com/golang/sys",
		toDir: "$GOPATH/src/golang.org/x/sys",
	}

	net := g{
		url:   "https://github.com/golang/net",
		toDir: "$GOPATH/src/golang.org/x/net",
	}

	content := []g{tools, sys, net}
	download(content)
	makeinstall()
}
