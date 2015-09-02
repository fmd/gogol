package engine

import (
    "github.com/go-gl/gl/v2.1/gl"
    "github.com/veandco/go-sdl2/sdl"
)

type Window struct {
    Title string
    Width int
    Height int
    Window *sdl.Window
    Context sdl.GLContext
}

type WindowOpts struct {
    Title string
    Width int
    Height int
}

func NewWindow(opts WindowOpts) *Window {
    var err error

    w := &Window{
        Title: opts.Title,
        Width: opts.Width,
        Height: opts.Height,
    }

    if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
        panic(err)
    }

    if err = gl.Init(); err != nil {
        panic(err)
    }

    w.Window, err = sdl.CreateWindow(w.Title,
                                     sdl.WINDOWPOS_UNDEFINED,
                                     sdl.WINDOWPOS_UNDEFINED,
                                     w.Width, w.Height, sdl.WINDOW_OPENGL)
    if err != nil {
        panic(err)
    }

    w.Context, err = sdl.GL_CreateContext(w.Window)
    if err != nil {
        panic(err)
    }

    gl.Enable(gl.DEPTH_TEST)
    gl.ClearColor(0.2, 0.2, 0.3, 1.0)
    gl.ClearDepth(1)
    gl.DepthFunc(gl.LEQUAL)
    gl.Viewport(0, 0, int32(w.Width), int32(w.Height))

    return w
}

func (w *Window) Swap() {
    sdl.GL_SwapWindow(w.Window)
}

func (w *Window) Clear() {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w *Window) Destroy() {
    sdl.GL_DeleteContext(w.Context)
    w.Window.Destroy()
    sdl.Quit()
}

