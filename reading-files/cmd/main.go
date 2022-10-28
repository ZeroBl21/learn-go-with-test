package main

import (
	"log"
	"os"

	blogposts "github.com/zerobl21/withTest/reading-files"
)

func main() {
  posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))

  if err != nil {
    log.Fatal(err)
  }

  log.Println(posts)
}
