package main

import (
	"todoapp/pkg/util/env"
	cmd "todoapp/src/command"
)

func init() {
	env := env.NewEnv()
	env.Load()
}

func main() {
  cmd.Init()
}