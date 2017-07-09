package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func parseDictionary(file *os.File) []string {
  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
     lines = append(lines, string(scanner.Bytes()))
  }

  return lines
}

func findWords(text string, words []string) []string {
  sameWord := true
  foundWords := make([]string, 0)
  for _, entry := range words {
    bEntry := []byte(entry)
    for i := range bEntry {
        if !strings.Contains(text, strings.ToLower(string(bEntry[i]))) {
          sameWord = false
          break
        }
    }
    if sameWord {
      foundWords = append(foundWords, entry)
    }
    sameWord = true
  }

  return foundWords
}

func main() {
  text := os.Args[2]
  file, err := os.Open(os.Args[1])
  checkErr(err)

  words := parseDictionary(file)
  results := findWords(text, words)
  for i := range results {
    fmt.Println(results[i])
  }
}
