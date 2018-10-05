package main

import (
	"io"
	"os"
	"fmt"
	"github.com/mgutz/ansi"
)

var writer io.Writer

func InitWriter()  {
	writer = io.MultiWriter(os.Stdout)
}

func Output(value map[string]Results)  {
	fmt.Fprintf(writer, "\n")

	var (
		ok = ansi.Color("Ok\n", "green")
		fail = ansi.Color("Fail\n", "red+b")
		okMark = ansi.Color("✓️\n", "green")
		failMark = ansi.Color("x\n", "red+b")
	)

	for suit, results := range value {
		fmt.Fprintf(writer, "SUIT: %s\n", ansi.Color(" " + suit + " ", "white:blue"))
		fmt.Fprintf(writer, "=================================================\n")

		count := 0

		for _, result := range results.List {
			if result.DataPass && result.ResPass {
				fmt.Fprintf(writer, "%-30s%24s", result.Title, ok)
			} else {
				fmt.Fprintf(writer, "%-33s%24s", result.Title, fail)
			}

			fmt.Fprintf(writer, "-------------------------------------------------\n")

			if result.ResPass {
				fmt.Fprintf(writer, "%-28s%24s", "响应比对", okMark)
			} else {
				fmt.Fprintf(writer, "%-29s%24s", "响应比对", failMark)
			}

			if result.DataPass {
				fmt.Fprintf(writer, "%-28s%24s", "数据比对", okMark)
			} else {
				fmt.Fprintf(writer, "%-29s%24s", "数据比对", failMark)
			}

			if !result.ResPass {
				fmt.Fprintf(writer, "-------------------------------------------------\n")
				fmt.Fprintf(writer, ansi.Color(" actual response \n\n", "yellow+b") + result.Description + "\n")
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
}
