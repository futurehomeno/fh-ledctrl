package main

import (
	"fmt"
	"time"

	ledctrl "github.com/futurehomeno/fh-ledctrl"
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
	initalColor, err := ledctrl.GetColor()
	if err != nil {
		panic(err)
	}
	for _, c := range colorList {
		fmt.Printf("Setting color %+v\n", c)
		err = ledctrl.SetColor(c)
		if err != nil {
			panic(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	err = ledctrl.SetColor(initalColor)
	if err != nil {
		panic(err)
	}
}
