package main

import (
	"./externals/github"
	"time"
)

func main() {
	github.Get("Nicovideo", "VideoCollection", 5)
	time.Sleep(10000 * time.Millisecond)
}
