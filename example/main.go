package main

import (
	"fmt"
	"time"

        "github.com/aditya43/clear-shell-screen-golang"
)

func main() {
	screen.Clear()

	for {
		screen.MoveTopLeft()

		w, h := screen.Size()
		fmt.Printf("Width: %d Height: %d\n", w, h)

		fmt.Println(time.Now())

		time.Sleep(time.Second)
	}
}
