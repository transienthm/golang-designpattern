package main

import (
	"strings"
	"fmt"
)

// Subject
type Git interface {
	Clone(url string) bool
}

//realSubject
type Github struct {

}

func (p Github) Clone(url string) bool {
	if strings.HasPrefix(url, "https") {
		fmt.Println("clone from " + url)
		return true
	}

	fmt.Println("failed to clone from " + url)
	return false
}

//proxy
type GitBash struct {
	GitCmd Git
}

func (p GitBash) Clone(url string) bool {
	return p.GitCmd.Clone(url)
}

func GetGit() Git {
	return GitBash{GitCmd:Github{}}
}

type Coder struct {
}

func (p Coder) GetCode(url string) {
	gitBash := GetGit()

	if gitBash.Clone(url) {
		fmt.Println("success...")
	} else {
		fmt.Println("failed...")
	}
}

func main() {
	coder := Coder{}

	coder.GetCode("https://www.baidu.com")
}