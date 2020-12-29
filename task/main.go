package main
import (
	"./cmd"
	"./db"
)

func main() {
	db.Init()
	cmd.Execute()
}
