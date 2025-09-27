package editor

import (
	"os"

	"github.com/mattiaizzi/light_editor/io"
)

type Editor struct {
	renderer io.Renderer
	content  []byte
}

func InitEditor(r io.Renderer, f string) (Editor, error) {
	editor := Editor{
		renderer: r,
	}
	e := editor.clear()
	if len(f) > 0 {
		e = editor.open(f)
	}
	return editor, e
}

func (e Editor) Render() error {
	return e.renderer.Render(e.content)
}

func (e *Editor) setContent(c []byte) {
	e.content = c
}

func (e *Editor) open(p string) error {
	f, err := os.ReadFile(p)
	if err != nil {
		return err
	}
	e.setContent(f)
	return nil
}

func (e Editor) clear() error {
	return e.renderer.Clear()
}
