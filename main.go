package main

import (
	"./externals/github"
	"fmt"
)

func main() {
	comments, _ := github.GetPullComments("Nicovideo", "nicovideo-web", 354)
	fmt.Println(comments)
}
