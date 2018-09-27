package main

import (
	"flag"
)

func main()  {

	var entrance string

	flag.StringVar(&entrance, "tests", "", "entrance json file")
	flag.Parse()

	if entrance == "" {
		panic("wrong parameter")
	}

	// TODO: 初始化database, io
	InitWriter()

	eng := New(entrance)
	eng.Run()
}
