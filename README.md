# ledctrl
A Go Library to easily control the hub's LED.

## Usage

### Setup
First, you need to make sure that your user is part of the `gpio` group. If you are developing a service, add this to your `postinst` file:

```shell script
usermod -a -G gpio myuser
```

Get the library

```shell script
go get github.com/futurehomeno/fh-ledctrl
```

Import the package

```go
import ledctrl "github.com/futurehomeno/fh-ledctrl"
```

### API
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

### Example
A usage example can be found in `example/main.go`.

## Hub Support
The library supports hub model `cube-1v1-eu` and newer.

To check the model of your hub, run:

```shell script
cat /usr/local/share/futurehome/smarthub.model
```

On model `cube-1v1-eu` and newer, fh-selftest makes sure that the LED GPIO pins are configured properly according to this script https://github.com/futurehomeno/fh-selftest/blob/develop/no-test/share/set-leds-white. For older hubs (``cube-1v0-eu``) and beta hubs (`cube-1v1-eu-proto-*`), the pins are not exported and thus those versions are not supported. An error will be returned when attempting to use the library on those hubs.

To use the library on unsupported hubs, you have to put the following:

```shell script
SUBSYSTEM=="gpio", KERNEL=="gpiochip*", ACTION=="add", PROGRAM="/bin/sh -c 'chown root:gpio /sys/class/gpio/export /sys/class/gpio/unexport ; chmod 220 /sys/class/gpio/export /sys/class/gpio/unexport'"
SUBSYSTEM=="gpio", KERNEL=="gpio*", ACTION=="add", PROGRAM="/bin/sh -c 'chown root:gpio /sys%p/active_low /sys%p/direction /sys%p/edge /sys%p/value ; chmod 660 /sys%p/active_low /sys%p/direction /sys%p/edge /sys%p/value'"
```

in a new file `/lib/udev/rules.d/60-gpio.rules` and reboot the hub. Then, export the pins by running:

```shell script
echo 11 > /sys/class/gpio/export
echo 12 > /sys/class/gpio/export
echo 13 > /sys/class/gpio/export
```