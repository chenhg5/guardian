package main

import (
	"io"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/mattn/go-runewidth"
)

var (
	writer    io.Writer
	finalPass = true
	ok        = ansi.Color("Ok\n", "green")
	pass      = ansi.Color("Pass\n", "green")
	fail      = ansi.Color("Fail\n", "red+b")
	okMark    = ansi.Color("✓️\n", "green")
	failMark  = ansi.Color("x\n", "red+b")
)

func InitWriter(w io.Writer) {
	writer = w
}

func LogResult(DataPass, ResPass, debug bool, SqlDesc, ResDesc, Title string) {

	if DataPass && ResPass {
		fmt.Fprintf(writer, "%s%24s", runewidth.FillRight(Title, 32), ok)
	} else {
		fmt.Fprintf(writer, "%s%24s", runewidth.FillRight(Title, 35), fail)
	}

	line()

	if ResPass {
		fmt.Fprintf(writer, "%-28s%24s", "响应比对", okMark)
	} else {
		finalPass = false
		fmt.Fprintf(writer, "%-29s%24s", "响应比对", failMark)
	}

	if DataPass {
		fmt.Fprintf(writer, "%-28s%24s", "数据比对", okMark)
	} else {
		finalPass = false
		fmt.Fprintf(writer, "%-29s%24s", "数据比对", failMark)
	}

	if !ResPass || debug {
		line()
		fmt.Fprintf(writer, ansi.Color("response \n\n", "yellow+b")+ResDesc+"\n")
	}

	if !DataPass || debug {
		line()
		fmt.Fprintf(writer, ansi.Color("sql \n\n", "yellow+b")+SqlDesc+"\n")
	}

	equalLine()
}

func LogTitle(suit string) {
	fmt.Fprintf(writer, "SUIT: %s\n", ansi.Color(" "+suit+" ", "white:blue"))
	equalLine()
}

func line() {
	fmt.Fprintf(writer, "-------------------------------------------------\n")
}

func equalLine() {
	fmt.Fprintf(writer, "=================================================\n")
}

func LogFinal() {

	fmt.Fprint(writer, "\n")

	if !finalPass {
		panic("测试没通过！")
	} else {
		fmt.Fprintf(writer, "%s", pass)
		fmt.Fprint(writer, "\n")
	}
}

func LogFirst() {
	fmt.Fprint(writer, "\n")
}
