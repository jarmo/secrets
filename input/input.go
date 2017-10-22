package input

import (
  "os"
  "fmt"
  "strings"
  "bufio"
)

func Ask(message string) string {
  fmt.Print(message)
  var value string
  fmt.Scanln(&value)
  return value
}

func AskMultiline(message string) string {
  fmt.Print(message)
  var value []string
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
      value = append(value, scanner.Text())
  }

  return strings.Join(value, "\n")
}
