package main

import "github.com/nsf/termbox-go"
import "fmt"

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
func printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}

func draw_container() {
	termbox.SetCell(0, 0, 0x250C, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(79, 0, 0x2510, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(0, 23, 0x2514, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(79, 23, 0x2518, termbox.ColorWhite, termbox.ColorBlack)

	for i := 1; i < 79; i++ {
		termbox.SetCell(i, 0, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, 23, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, 17, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, 3, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
	}

	for i := 1; i < 23; i++ {
		termbox.SetCell(0, i, 0x2502, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(79, i, 0x2502, termbox.ColorWhite, termbox.ColorBlack)
	}

	termbox.SetCell(0, 17, 0x251C, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(79, 17, 0x2524, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(0, 3, 0x251C, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(79, 3, 0x2524, termbox.ColorWhite, termbox.ColorBlack)

	// Code for Blue band on the side.
	for i := 4; i < 17; i++ {
		termbox.SetCell(1, i, 0x2588, termbox.ColorBlue, termbox.ColorBlue)
		termbox.SetCell(78, i, 0x2588, termbox.ColorBlue, termbox.ColorBlue)
	}
	printf_tb(33, 1, termbox.ColorMagenta|termbox.AttrBold, termbox.ColorBlack, "Synapse RF")
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	draw_container()
	termbox.Flush()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			}
		}
	}
}
