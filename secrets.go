package main

import (
  "os"
  "github.com/jarmo/secrets/cli"
)

const VERSION = "3.1.0"

func main() {
  cli.Command(VERSION, os.Args[1:]).Execute()
}

