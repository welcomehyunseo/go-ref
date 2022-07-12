package main

import (
	"time"

	"atomicgo.dev/cursor"
)

func main() {
	area := cursor.NewArea()
	area.Update("hello!")
	time.Sleep(2 * time.Second)
	area.Clear()
	area.Update("bye!")
}
