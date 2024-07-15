package main

import (
	"github.com/fffunky/ansi-lib/ansi"
)

func main() {
	my_style := ansi.NewStyle(ansi.BOLD, ansi.ITALICS, ansi.FG_CYAN, ansi.BG_GREEN)
	my_style2 := ansi.NewStyle(ansi.BOLD, ansi.FG_RED)
	my_style3 := ansi.NewStyle(ansi.DIM, ansi.FG_BLUE, ansi.BG_WHITE)

	ansi.Aprint(my_style, "this function acts like fmt.Print\n")
	ansi.Aprintln(my_style2, "this function acts like fmt.Println")
	ansi.Aprintf(my_style3, "this function acts like %s\n", "fmt.Printf")
}
