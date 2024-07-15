# ansi-lib

A library that wraps many useful ansi escape codes for special terminal printing and formatting.

## Making a style

Use the function `NewStyle(args ...string) *Style`, passing in the library constants as args.
ex:
```go
import "github.com/fffunky/ansi-lib/ansi"

// If we want our style to be bold and italics, with magenta text and a yellow
// background, we do it like this.
var my_style *Style = NewStyle(ansi.BOLD, ansi.ITALICS, ansi.FG_MAGENTA, ansi.BG_YELLOW)
```

## Using the style

You may use any of the Printf analogs implemented in the library to print out text with these styles
```go
Aprint(style *Style, msg string) {...}
Aprintln(style *Style, msg string) {...}
Aprintf(style *Style, msg string, args ...any) {...}
```

for example:
```go
func main() {
	my_style := ansi.NewStyle(ansi.BOLD, ansi.ITALICS, ansi.FG_CYAN, ansi.BG_GREEN)
	my_style2 := ansi.NewStyle(ansi.BOLD, ansi.FG_RED)
	my_style3 := ansi.NewStyle(ansi.DIM, ansi.FG_BLUE, ansi.BG_WHITE)

	ansi.Aprint(my_style, "this function acts like fmt.Print\n")
	ansi.Aprintln(my_style2, "this function acts like fmt.Println")
	ansi.Aprintf(my_style3, "this function acts like %s\n", "fmt.Printf")
}
```
creates the following output:
![Aprint, Aprintln, Aprintf examples implemented with 3 different styles](/images/aprint_example.png)
