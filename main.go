package main

import (
	"flag"
	"io"
	"os"
)

func main()  {

	var entrance string

	flag.StringVar(&entrance, "tests", "", "entrance json file")
	flag.Parse()

	if entrance == "" {
		panic("wrong parameter")
	}

	InitWriter(io.MultiWriter(os.Stdout))

	New(entrance).Run()
}
