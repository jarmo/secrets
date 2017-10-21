package main

import (
  "os"
  "fmt"
  "github.com/jarmo/secrets/cli"
  "github.com/jarmo/secrets/commands"
  "github.com/jarmo/secrets/vault"
)

const VERSION = "0.0.1"

func main() {
  switch command := cli.Execute(VERSION, os.Args[1:]).(type) {
    case commands.List:
      fmt.Println(vault.List(command.Filter))
    case commands.Add:
      fmt.Println("Added:", vault.Add(command.Name))
    default:
      fmt.Printf("Unhandled command: %T\n", command)
  }
}
