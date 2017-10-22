package main

import (
  "os"
  "fmt"
  "github.com/jarmo/secrets/cli"
  "github.com/jarmo/secrets/command"
  "github.com/jarmo/secrets/vault"
)

const VERSION = "0.0.1"

func main() {
  switch parsedCommand := cli.Execute(VERSION, os.Args[1:]).(type) {
    case command.List:
      secrets := vault.List(parsedCommand.Filter)
      for _, secret := range secrets {
        fmt.Println(secret)
      }
    case command.Add:
      fmt.Println("Added:", vault.Add(parsedCommand.Name))
    default:
      fmt.Printf("Unhandled command: %T\n", parsedCommand)
  }
}
