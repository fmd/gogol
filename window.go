package bronson

import (
    "github.com/go-gl/gl/v2.1/gl"
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
)

type Window struct {
    Title string
    Width int32
    Height int32
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
        Width: int32(opts.Width),
        Height: int32(opts.Height),
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
                                     int(w.Width), int(w.Height), sdl.WINDOW_OPENGL)
    if err != nil {
        panic(err)
    }

    w.Context, err = sdl.GL_CreateContext(w.Window)
    if err != nil {
        panic(err)
    }

    /*
    gl.Enable(gl.DEPTH_TEST)
    gl.ClearColor(0.2, 0.2, 0.3, 1.0)
    gl.ClearDepth(1)
    gl.DepthFunc(gl.LEQUAL)
    gl.Viewport(0, 0, int32(w.Width), int32(w.Height))
    */

    w.SetVideoOptions()

    return w
}

func (w *Window) SetVideoOptions() {
    fmt.Println(fmt.Sprintf(" -Vendor: %s", gl.GetString(gl.VENDOR)))
    fmt.Println(fmt.Sprintf(" -Renderer: %s\n", gl.GetString(gl.RENDERER)))
    fmt.Println(fmt.Sprintf(" -Version: %s\n", gl.GetString(gl.VERSION)))
    fmt.Println(fmt.Sprintf(" -Extensions: %s\n", gl.GetString(gl.EXTENSIONS)))

    gl.ClearColor(0.0,0.0,0.0,1.0)

    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()

    gl.Viewport(0, 0, w.Width, w.Height)
    gl.Ortho(-float64(w.Width)/2, float64(w.Width)/2, -float64(w.Height)/2, float64(w.Height)/2, -1.0, 1.0)

    //Required for per-pixel placing.
    gl.Translatef (0.375, 0.375, 0.0)

    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()

    gl.Enable(gl.TEXTURE_2D)
    gl.Disable(gl.LIGHTING)
    gl.Disable(gl.DITHER)

    gl.Enable(gl.CULL_FACE)
    gl.CullFace(gl.FRONT)

    gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

    gl.Enable(gl.BLEND)
    gl.Enable(gl.ALPHA_TEST)
    gl.DepthFunc(gl.LEQUAL)

    gl.Disable(gl.DEPTH_TEST)
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

