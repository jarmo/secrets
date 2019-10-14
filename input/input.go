package input

import (
  "fmt"
  "syscall"
  "golang.org/x/crypto/ssh/terminal"
)

func AskPassword(message string) []byte {
  fmt.Print(message)
  password, err := terminal.ReadPassword(int(syscall.Stdin))
  if err != nil {
    panic(err)
  }
  fmt.Println()

  return password
}
