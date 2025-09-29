package io

import (
	"io"
	"os"
)

type Renderer interface {
	Render(content []byte) error
	Clear() error
}

type InputHandler interface {
	Read() error
}

type ANSIRenderer struct {
	writer io.Writer
}

func InitANSIRenderer() ANSIRenderer {
	return ANSIRenderer{writer: os.Stdout}
}

func (r ANSIRenderer) Render(content []byte) error {
	_, err := r.writer.Write(content)
	return err
}

func (r ANSIRenderer) Clear() error {
	/**
	*	\033 escape
	*	[ csi
	*	2J clear terminal
	*	H move cursor to home(0,0)
	**/
	command := "\033[2J\033[H"
	return r.Render([]byte(command))
}
