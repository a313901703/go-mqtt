package main

import (
	"mqtt/cmd"
	_ "mqtt/task"
)

func main() {
	cmd.Start()
}
