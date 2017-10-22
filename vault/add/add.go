package add

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "github.com/jarmo/secrets/secret"
)

func Execute(secrets []secret.Secret, name string) (secret.Secret, []secret.Secret) {
  fmt.Printf("Enter value for '%s':\n", name)
  var value []string
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
      value = append(value, scanner.Text())
  }

  newSecret := secret.Create(name, strings.Join(value, "\n"))
  return newSecret, append(secrets, newSecret)
}
