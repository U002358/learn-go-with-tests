package main

import (
	"fmt"
	blogposts "learn-go-with-tests/17-reading-files"
	"log"
	"os"
)

func main() {
	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)
	posts, err := blogposts.NewPostsFromFS(os.DirFS(currentDir + "\\cmd\\posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
