package io

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)

type Renderer interface {
	MoveCursor(row uint, col uint) error
	Render(content []byte) error
	Clear() error
}

type Cursor struct {
	row uint
	col uint
}

type InputHandler interface {
	Read() ([]byte, error)
}

type ANSIRenderer struct {
	cursor Cursor
	writer io.Writer
}

type TerminalInputHandler struct {
	reader io.Reader
}

func InitTerminalInputHandler() TerminalInputHandler {
	return TerminalInputHandler{
		reader: os.Stdin,
	}
}

func (t TerminalInputHandler) Read() ([]byte, error) {
	buf := make([]byte, 3)
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	_, err = t.reader.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil

}

func InitANSIRenderer() ANSIRenderer {
	return ANSIRenderer{
		writer: os.Stdout,
		cursor: Cursor{row: 0, col: 0},
	}
}

func (r ANSIRenderer) Render(content []byte) error {
	_, err := r.writer.Write(content)
	return err
}

func (r ANSIRenderer) MoveCursor(row uint, col uint) error {
	r.cursor.col = col
	r.cursor.row = row
	command := fmt.Sprintf("\033[%v;%vH", row, col)
	_, err := r.writer.Write([]byte(command))
	return err
}

func (r ANSIRenderer) Clear() error {
	/**
	*	\033 escape
	*	[ csi
	*	2J clear terminal
	**/
	command := "\033[2J"
	_, err := r.writer.Write([]byte(command))
	r.MoveCursor(0, 0)
	return err
}
