package ledctrl

import (
	"fmt"
	"os"
)

const (
	redPin   = 11
	greenPin = 12
	bluePin  = 13
)

type color struct {
	red   int
	green int
	blue  int
}

var (
	Red     = color{1, 0, 0}
	Green   = color{0, 1, 0}
	Blue    = color{0, 0, 1}
	White   = color{1, 1, 1}
	Off     = color{0, 0, 0}
	Yellow  = color{1, 1, 0}
	Cyan    = color{0, 1, 1}
	Magenta = color{1, 0, 1}
)

func writeToGPIO(value int, pin int) error {
	if value != 0 && value != 1 {
		return fmt.Errorf("Value %d must be either 0 or 1", value)
	}
	filename := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(fmt.Sprintf("%d", value)))
	return err
}

func SetColor(c color) error {
	var err error
	err = writeToGPIO(c.red, redPin)
	if err != nil {
		return err
	}
	err = writeToGPIO(c.green, greenPin)
	if err != nil {
		return err
	}
	err = writeToGPIO(c.blue, bluePin)
	if err != nil {
		return err
	}
	return nil
}
