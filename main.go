package main

import (
  "fmt"
  "os"
  "commands"
)

func main() {

  if err := commands.OctopusCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

}