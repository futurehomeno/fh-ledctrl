package ledctrl

import "testing"

var colorList = []color{
	Red, Green, Blue, White, Off, Yellow, Cyan, Magenta,
}

func TestSetColor(t *testing.T) {
	var err error
	for _, c := range colorList {
		err = SetColor(c)
		if err != nil {
			t.Errorf("Got an unexpected error: %s", err.Error())
		}
	}
}
