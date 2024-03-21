package main

import (
	"server/cmd"
	_ "server/internal/db"
)

func main() {
	cmd.Execute()
}
