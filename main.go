package main

import (
	"fmt"

	"github.com/mattiaizzi/light_editor/editor"
	"github.com/mattiaizzi/light_editor/io"
)

func main() {
	r := io.InitANSIRenderer()
	e, err := editor.InitEditor(r, "file.txt")
	if err != nil {
		fmt.Println(err)
	}
	e.Render()
}
