package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "os"
  "regexp"
)

type FileSystemNode struct {
  Name     string           `json:"name"`
  IsFile   bool             `json:"isFile"`
  Children []FileSystemNode `json:"children"`
}

var hackedFileRE = regexp.MustCompile(`\.hack$`)

func (node *FileSystemNode) countHackedFiles() int {
  count := 0
  if node.IsFile && hackedFileRE.MatchString(node.Name) {
    return 1
  }
  for _, child := range node.Children {
    count += child.countHackedFiles()
  }
  return count
}

func main() {
  in := bufio.NewReader(os.Stdin)
  out := bufio.NewWriter(os.Stdout)
  defer out.Flush()

  var numTests int
  fmt.Fscanln(in, &numTests)

  for i := 0; i < numTests; i++ {
    var numLines int
    fmt.Fscanln(in, &numLines)

    data := []byte{}
    for j := 0; j < numLines; j++ {
      line, _ := in.ReadBytes('\n')
      data = append(data, line...)
    }

    var root FileSystemNode
    if err := json.Unmarshal(data, &root); err != nil {
      fmt.Println("Error unmarshalling JSON:", err)
      continue
    }

    fmt.Println(root.countHackedFiles())
  }
}