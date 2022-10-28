package blogposts

import "io/fs"

type Post struct{
  Title string
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
  postFile, err := fileSystem.Open(fileName)

  if err != nil {
    return Post{}, err
  }

  defer postFile.Close()

  return newPost(postFile)
}
