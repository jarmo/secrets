package main

import (
  "os"
  "fmt"
  "github.com/jarmo/secrets/cli"
)

const VERSION = "0.0.1"

func main() {
  switch command := cli.Execute(VERSION, os.Args[1:]).(type) {
    default:
      fmt.Printf("Unhandled command: %T\n", command)
  }
}
