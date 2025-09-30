package main

import (
	"fmt"

	"github.com/mattiaizzi/light_editor/editor"
	"github.com/mattiaizzi/light_editor/io"
)

func main() {
	//	args := os.Args[1:]
	//	p := args[0]
	r := io.InitANSIRenderer()
	e := editor.InitEditor(r)
	e.Clear()
	// err := e.Open(p)
	handler := io.InitTerminalInputHandler()
	for {
		buf, err := handler.Read()
		if err != nil {
			panic(1)
		}
		c := buf[0]
		if c >= 32 && c < 127 {
			fmt.Println(string(c))
		}
		if c == 27 {
			panic(1)
		}
	}
	//	if err != nil {
	//		log.Fatalf("Error: %v", err)
	//	}
	//
	// e.Render()
}
