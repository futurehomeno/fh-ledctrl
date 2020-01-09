package ledctrl

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	redPin   = 11
	greenPin = 12
	bluePin  = 13
)

var pins = []int{redPin, greenPin, bluePin}

const (
	gpioExportFile    = "/sys/class/gpio/export"
	gpioValueFile     = "/sys/class/gpio/gpio%d/value"
	gpioDirectionFile = "/sys/class/gpio/gpio%d/direction"
)

type Color struct {
	red   int
	green int
	blue  int
}

var (
	Red     = Color{1, 0, 0}
	Green   = Color{0, 1, 0}
	Blue    = Color{0, 0, 1}
	White   = Color{1, 1, 1}
	Off     = Color{0, 0, 0}
	Yellow  = Color{1, 1, 0}
	Cyan    = Color{0, 1, 1}
	Magenta = Color{1, 0, 1}
)

// writeToGPIO writes the passed value (0 or 1) to the specified GPIO pin number
// The function assumes the pin has been exported and set as output,
// which is the case on CubeOS image
func writeToGPIO(value int, pin int) error {
	if value != 0 && value != 1 {
		return fmt.Errorf("Value %d must be either 0 or 1", value)
	}
	filename := fmt.Sprintf(gpioValueFile, pin)
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(fmt.Sprintf("%d", value)))
	return err
}

func readGPIO(pin int) (int, error) {
	filename := fmt.Sprintf(gpioValueFile, pin)
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return -1, err
	}
	bytes := make([]byte, 1)
	_, err = file.Read(bytes)
	if err != nil {
		return -1, err
	}
	bit, err := strconv.Atoi(string(bytes))
	if err != nil {
		return -1, err
	}
	if bit != 0 && bit != 1 {
		return -1, fmt.Errorf("Pin value must be either 0 or 1. Instead, got %d", bit)
	}
	return bit, nil
}

// SetColor changes the hub LED's color to the passed one
func SetColor(c Color) error {
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

// GetColor returns the current color of the LED
func GetColor() (Color, error) {
	var err error
	red, err := readGPIO(redPin)
	if err != nil {
		return Off, err
	}
	green, err := readGPIO(greenPin)
	if err != nil {
		return Off, err
	}
	blue, err := readGPIO(bluePin)
	if err != nil {
		return Off, err
	}
	return Color{red, green, blue}, nil
}

// If the hub is an older model, export the LED GPIO pins
func init() {
	const hubModelFile = "/usr/local/share/futurehome/smarthub.model"
	hubModel, err := ioutil.ReadFile(hubModelFile)
	if err != nil {
		return
	}
	if matched, _ := regexp.MatchString(`cube-1v1-eu-proto-\d`, string(hubModel)); matched {
		expFile, err := os.OpenFile(gpioExportFile, os.O_WRONLY, 0666)
		if err != nil {
			return
		}
		defer expFile.Close()
		for _, pin := range pins {
			_, err = expFile.Write([]byte(strconv.Itoa(pin)))
			if err != nil {
				return
			}
		}
		time.Sleep(3 * time.Second)
		for _, pin := range pins {
			filename := fmt.Sprintf(gpioDirectionFile, pin)
			dirFile, err := os.OpenFile(filename, os.O_WRONLY, 0666)
			if err != nil {
				return
			}
			_, err = dirFile.Write([]byte("out"))
			if err != nil {
				return
			}
			dirFile.Close()
		}
	}
}
