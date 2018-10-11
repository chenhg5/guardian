package main

import (
	"io"
	"os"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/mattn/go-runewidth"
)

var writer io.Writer

func InitWriter()  {
	writer = io.MultiWriter(os.Stdout)
}

func Output(value map[string]Results, debug bool)  {
	fmt.Fprintf(writer, "\n")

	var (
		ok = ansi.Color("Ok\n", "green")
		fail = ansi.Color("Fail\n", "red+b")
		okMark = ansi.Color("✓️\n", "green")
		failMark = ansi.Color("x\n", "red+b")
		finalPass = true
	)

	for suit, results := range value {
		fmt.Fprintf(writer, "SUIT: %s\n", ansi.Color(" " + suit + " ", "white:blue"))
		fmt.Fprintf(writer, "=================================================\n")

		count := 0

		for _, result := range results.List {
			if result.DataPass && result.ResPass {
				fmt.Fprintf(writer, "%s%24s", runewidth.FillRight(result.Title, 32), ok)
			} else {
				fmt.Fprintf(writer, "%s%24s", runewidth.FillRight(result.Title, 35), fail)
			}

			fmt.Fprintf(writer, "-------------------------------------------------\n")

			if result.ResPass {
				fmt.Fprintf(writer, "%-28s%24s", "响应比对", okMark)
			} else {
				finalPass = false
				fmt.Fprintf(writer, "%-29s%24s", "响应比对", failMark)
			}

			if result.DataPass {
				fmt.Fprintf(writer, "%-28s%24s", "数据比对", okMark)
			} else {
				finalPass = false
				fmt.Fprintf(writer, "%-29s%24s", "数据比对", failMark)
			}

			if !result.ResPass || debug {
				fmt.Fprintf(writer, "-------------------------------------------------\n")
				fmt.Fprintf(writer, ansi.Color("response \n\n", "yellow+b") + result.ResDesc + "\n")
			}

			if !result.DataPass || debug {
				fmt.Fprintf(writer, "-------------------------------------------------\n")
				fmt.Fprintf(writer, ansi.Color("sql \n\n", "yellow+b") + result.SqlDesc + "\n")
			}

			fmt.Fprintf(writer, "=================================================\n")
			count++
		}

		fmt.Fprint(writer, "\n")

		if results.Pass {
			fmt.Fprintf(writer, ok)
		} else {
			fmt.Fprintf(writer, fail)
		}

		fmt.Fprint(writer, "\n")
	}

	if !finalPass {
		panic("测试没通过！")
	}
}
