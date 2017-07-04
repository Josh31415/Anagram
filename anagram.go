package main

import (
  "fmt"
  "os"
  "bufio"
)

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func findWords() {

}

func parseDictionary(file *os.File, letter byte, words chan []string) {
  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if scanner.Bytes()[0] == letter {
      lines = append(lines, string(scanner.Bytes()))
    }
  }

  words <- lines
}

func main() {
  text := os.Args[1]
  file, err := os.Open("dictionary.txt")
  fmt.Println(file)
  checkErr(err)
  fileParts := make([]chan []string, 26)
  fileSections := make([][]string, 26)

  for i := range fileParts {
    char := i + 97
    go parseDictionary(file, byte(char), fileParts[i])
  }

  for i := range fileParts {
    fileSections[i] = <-fileParts[i]
    fmt.Print(fileSections[i])
  }

 letters := []byte(text)
 fmt.Print(letters)
}
