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
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  return scanner.Text()
}

func AskMultiline(message string) string {
  fmt.Print(message)
  var value []string
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
      value = append(value, scanner.Text())
  }

  return replaceUnprintableCharacters(strings.Join(value, "\n"))
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

func replaceUnprintableCharacters(s string) string {
  ctrlD, ctrlX, ctrlZ := "\x04", "\x18", "\x1A"

  return strings.Replace(
    strings.Replace(
      strings.Replace(s,
      ctrlD, "", -1),
      ctrlX, "", -1),
      ctrlZ, "", -1)
}
