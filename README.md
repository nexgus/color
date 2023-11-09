# color

This is a simplified version of [faith/color](https://github.com/fatih/color). To me, the original usage of [faith/color](https://github.com/fatih/color) is too hard to use. I need a simple, intuitivement module.

## Example
```go
import (
    "fmt"

    "github.com/nexgus/color"
)

func main() {
    fmt.Printf("%sThis is simple and intuitivement.%s\n", color.SetColor(color.FgHiWhite, color.BgRed), color.Reset())
}
```
