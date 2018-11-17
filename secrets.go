package main

import (
  "os"
  "github.com/jarmo/secrets/cli"
)

const VERSION = "4.0.0"

func main() {
  cli.Command(VERSION, os.Args[1:]).Execute()
}

