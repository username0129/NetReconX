package main

import (
	"server/cmd"
	_ "server/internal/database"
)

func main() {
	cmd.Execute()
}
