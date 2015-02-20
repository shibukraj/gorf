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

func print_mainmenu() {
	printf_tb(33, 4, termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault, "Main Menu")
	printf_tb(20, 6, termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault, "(01)-Receiving")
	printf_tb(20, 7, termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault, "(02)-Picking")
	printf_tb(20, 8, termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault, "(03)-Shipping")
	printf_tb(20, 9, termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault, "(04)-ASN Receiving")

	printf_tb(15, 12, termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault, "Enter Selection : __")
	//	termbox.SetCursor(33, 12)
	//	termbox.SetInputMode(termbox.InputEsc)
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

	print_mainmenu()

}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	draw_container()
	var edit_box EditBox
	draw_editbox(33, 12, 20)
	termbox.Flush()
loop:
	for {
		var selectedValue string
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowLeft, termbox.KeyCtrlB:
				edit_box.MoveCursorOneRuneBackward()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				edit_box.DeleteRuneBackward()
			case termbox.KeyEnter:
				selectedValue = string(edit_box.text[:])
			default:
				if ev.Ch != 0 {
					// Insert the character to the position in the screen.
					edit_box.InsertRune(ev.Ch)
				}
			}
			edit_box.Draw(33, 12, 20, 1)
			termbox.SetCursor(33+edit_box.CursorX(), 12)
			printf_tb(33, 14, termbox.ColorMagenta|termbox.AttrBold, termbox.ColorBlack, selectedValue)
			termbox.Flush()
		}
	}
}
