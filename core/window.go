package core

import "github.com/go-gl/glfw/v3.3/glfw"

type Window interface {
	SetSize(width int, height int)
	SetPosition(x int, y int)
	SetTitle(title string)
	ShouldClose() bool
	Destroy()
}

type WindowImp struct {
	handle *glfw.Window
}

func NewWindow(width int, height int, title string) Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	return &WindowImp{
		handle: window,
	}
}

func (w *WindowImp) SetSize(width int, height int) {
	if w.handle != nil {
		w.handle.SetSize(width, height)
	}
}

func (w *WindowImp) SetPosition(x int, y int) {
	if w.handle != nil {
		w.handle.SetPos(x, y)
	}
}

func (w *WindowImp) SetTitle(title string) {
	if w.handle != nil {
		w.handle.SetTitle(title)
	}
}

func (w *WindowImp) ShouldClose() bool {
	return w.ShouldClose()
}

func (w *WindowImp) Destroy() {
	if w.handle != nil {
		w.handle.Destroy()
	}

	glfw.Terminate()
}
