# fh-ledctrl
A Library to easily control the hub's LED

## Usage
Get the library

```shell script
go get github.com/futurehomeno/fh-ledctrl
```

The API has one function

```go
func SetColor(c Color) error
```

The function takes one of eight colors:

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
Basically, just do this:

```go
import ledctrl "github.com/futurehomeno/fh-ledctrl"

func main() {
    err = ledctrl.SetColor(ledctrl.Green)
    if err != nil {
        panic(err)
    }
}
```