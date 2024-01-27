package main

import (
	"github.com/apioo/typehub-cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cmd.Execute()
}
