# ledctrl
A Go Library to easily control the hub's LED.

The library works assuming fh-selftest has run and the LED GPIO pins are configured properly according to this script https://github.com/futurehomeno/fh-selftest/blob/develop/no-test/share/set-leds-white

## Usage
Get the library

```shell script
go get github.com/futurehomeno/fh-ledctrl
```

The API has the following functions

```
func SetColor(c Color) error
func GetColor() (Color, error)
```

The following colors are supported:

```
ledctrl.Red,
ledctrl.Green,
ledctrl.Blue,
ledctrl.White,
ledctrl.Off,
ledctrl.Yellow,
ledctrl.Cyan,
ledctrl.Magenta,
```

## Example
A usage example can be found in `example/main.go`.

First, import the package

```go
import ledctrl "github.com/futurehomeno/fh-ledctrl"
```

To set the LED color:

```go
err := ledctrl.SetColor(ledctrl.Green)
```

To get the LED color:

```go
color, err := ledctrl.GetColor()
```
