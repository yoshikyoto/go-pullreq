package main

import (
	"./externals/github"
	"./repos"
	"fmt"
)

func main() {
	comments, _ := github.GetPullComments("Nicovideo", "VideoCollection", 5)
	fmt.Println(comments)
	// 非同期でDBに突っ込みまくる
	for _, comment := range comments {
		go repos.Create(comment)
	}
}
