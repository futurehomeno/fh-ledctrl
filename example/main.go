package main

import (
	"fmt"
	ledctrl "github.com/futurehomeno/fh-ledctrl"
	"time"
)

var colorList = []ledctrl.Color{
	ledctrl.Red,
	ledctrl.Green,
	ledctrl.Blue,
	ledctrl.White,
	ledctrl.Off,
	ledctrl.Yellow,
	ledctrl.Cyan,
	ledctrl.Magenta,
}

func main() {
	var err error
	for _, c := range colorList {
		fmt.Printf("Setting color %+v\n", c)
		err = ledctrl.SetColor(c)
		if err != nil {
			fmt.Println("Got an unexpected error: %s", err.Error())
		}
		time.Sleep(500 * time.Millisecond)
	}
}
