package editor

import (
	"bytes"
	"os"

	"github.com/mattiaizzi/light_editor/io"
)

type Editor struct {
	renderer io.Renderer
	buffer   *Buffer
}

func InitEditor(r io.Renderer) Editor {
	editor := Editor{
		renderer: r,
		buffer:   new(Buffer),
	}
	return editor
}

func (e *Editor) Open(p string) error {
	f, err := os.ReadFile(p)
	if err != nil {
		return err
	}
	e.setContent(f)
	return nil
}

func (e Editor) Clear() error {
	return e.renderer.Clear()
}

func (e Editor) Render() error {
	content := make([][]byte, len(e.buffer.lines))
	for i, l := range e.buffer.lines {
		content[i] = make([]byte, len(l.content))
		copy(content[i], l.content)
	}
	return e.renderer.Render(bytes.Join(content, []byte("\n")))
}

func (e *Editor) setContent(c []byte) {
	e.buffer = new(Buffer)
	for l := range bytes.Lines(c) {
		bl := newLine(l[:len(l)-1])
		// TODO move add lines to buffer type
		e.buffer.lines = append(e.buffer.lines, bl)
	}
}

type Line struct {
	content []byte
}

type Buffer struct {
	lines []*Line
}

func newLine(content []byte) *Line {
	// Initialize with multiple of 64  as capacity to avoid reallocations
	diff := len(content) % 64
	line := Line{}
	line.content = append(make([]byte, 0, len(content)+64-diff), content...)
	return &line
}

func (l *Line) insertChar(index uint, char byte) {
	// Check if index is out of bound
	if index > uint(len(l.content)) {
		index = uint(len(l.content))
	}
	// Check if the slice could contain the character https://go.dev/wiki/SliceTricks
	if cap(l.content)-len(l.content) <= 0 {
		l = newLine(l.content)
	}
	// Append "nil" element
	l.content = append(l.content, 0)
	// From the index position shift right the content
	copy(l.content[index+1:], l.content[index:])
	// Replace the character
	l.content[index] = char
}
func (l *Line) deleteChar(index uint) {
	copy(l.content[index:], l.content[index+1:])
	l.content[len(l.content)-1] = 0
	l.content = l.content[:len(l.content)-1]
}
