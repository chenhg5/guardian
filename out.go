package guardian

import (
	"io"
	"os"
	"fmt"
)

var writer io.Writer

func InitWriter()  {
	writer = io.MultiWriter(os.Stdout)
}

func Output(value map[string]Results)  {
	for suit, results := range value {
		fmt.Fprintf(writer, "SUIT: %v", suit)
		fmt.Fprintf(writer, "=================================================\n")

		count := 0

		for _, result := range results.List {
			if count != 0 {
				fmt.Fprintf(writer, "=================================================\n")
			}
			if result.Pass {
				fmt.Fprintf(writer, "%v                                   Ok\n", result.Title)
			} else {
				fmt.Fprintf(writer, "%v                                   ERROR\n", result.Title)
			}

			fmt.Fprintf(writer, "=================================================\n")
			count++
		}

		fmt.Fprint(writer, "\n")

		if results.Pass {
			fmt.Fprintf(writer, "Ok")
		} else {
			fmt.Fprintf(writer, "Error")
		}
	}
}
