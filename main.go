package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type result struct {
	dir []string
}

func (ori *result) append(wd string) {
	ori.dir = append(ori.dir, wd)
}

//Just run "Go install"
func runInstall(wd string) {
	os.Chdir(wd)
	log.Println("Current make install in ", wd)
	cmd := exec.Command("go", "install")
	cmd.Run()
}

//Confirm if the dir need to run "Go install"
func containGo(wd string) (flag bool) {
	r, _ := ioutil.ReadDir(wd)
	for _, subItems := range r {
		if filepath.Ext(subItems.Name()) == ".go" {
			return true
		}
	}
	return false
}

//Find the dir needed to run "Go install"
func findDir(rootPath string, directiories *result) {
	var path string
	r, _ := ioutil.ReadDir(rootPath)
	if containGo(rootPath) {
		directiories.append(rootPath)
	} else {
		for _, i := range r {
			path = rootPath + i.Name() + "\\"
			findDir(path, directiories)
		}
	}

}
func main() {
	gopath := os.Getenv("GOPATH")
	rootpath := gopath + "\\src\\golang.org\\x\\"
	var directiories result //The needed directiories
	findDir(rootpath, &directiories)
	for _, dir := range directiories.dir {
		runInstall(dir)
	}
}
