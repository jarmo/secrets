package main

import (
  "os"
  "github.com/jarmo/secrets/v5/cli"
)

const VERSION = "5.0.0"

func main() {
  cli.Command(VERSION, os.Args[1:]).Execute()
}

