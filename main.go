package main

import (
	"log"
	"os"

	"github.com/mattiaizzi/light_editor/editor"
	"github.com/mattiaizzi/light_editor/io"
)

func main() {
	args := os.Args[1:]
	p := args[0]
	r := io.InitANSIRenderer()
	e := editor.InitEditor(r)
	err := e.Open(p)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	e.Render()
}
