package add

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "github.com/jarmo/secrets/secret"
)

func Execute(secrets []secret.Secret, name string) secret.Secret {
  fmt.Println("Enter value:")
  var value []string
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
      value = append(value, scanner.Text())
  }

  return secret.Create(name, strings.Join(value, "\n"))
}
