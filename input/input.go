package input

import (
  "os"
  "fmt"
  "strings"
  "bufio"
  "syscall"
  "golang.org/x/crypto/ssh/terminal"
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

func AskPassword(message string) []byte {
  fmt.Print(message)
  password, err := terminal.ReadPassword(int(syscall.Stdin))
  if err != nil {
    panic(err)
  }
  fmt.Println()

  return password
}
