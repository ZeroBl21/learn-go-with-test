package blogposts

import (
	"io"
	"io/fs"
)

// NewPostsFromFS returns a collection of blog posts from a file system. If it does not conform to the format then it'll return an error
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
  dir, err := fs.ReadDir(fileSystem, ".")

  if err != nil {
    return nil, err
  }

  var posts []Post

  for _, f := range dir {
    post, err := getPost(fileSystem, f.Name())

    if err != nil {
      return nil, err
    }

    posts = append(posts, post)
  }

	return posts, nil
}

func newPost(postFile io.Reader) (Post, error) {
  postData, err := io.ReadAll(postFile)

  if err != nil {
    return Post{}, err
  }

  post := Post{Title: string(postData)[7:]}

  return post, nil
}

