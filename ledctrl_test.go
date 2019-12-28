package ledctrl

import (
	"fmt"
	"testing"
	"time"
)

var colorList = []Color{
	Red, Green, Blue, White, Off, Yellow, Cyan, Magenta,
}

func TestSetColor(t *testing.T) {
	var err error
	for _, c := range colorList {
		fmt.Printf("Setting color %+v\n", c)
		err = SetColor(c)
		if err != nil {
			t.Errorf("Got an unexpected error: %s", err.Error())
		}
		time.Sleep(500 * time.Millisecond)
	}
}
