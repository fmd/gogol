package gogol

import (
    "github.com/go-gl/gl/v2.1/gl"
    "github.com/veandco/go-sdl2/sdl"
)

// The options to pass in when creating a new window.
type WindowOpts struct {
    Title string
    Width int
    Height int
}

// The Window type.
// There should only ever be one of these.
type Window struct {
    Title string
    Width int
    Height int
    Window *sdl.Window
    Context sdl.GLContext
}

// Creates a new Bronson window.
func NewWindow(opts WindowOpts) *Window {
    w := &Window{
        Title: opts.Title,
        Width: opts.Width,
        Height: opts.Height,
    }

    initSDL()
    initGL()

    w.createWindow()
    w.createContext()
    w.setVideoOptions()

    return w
}

// Swaps the OpenGL window.
// Call this after everything has been rendered.
func (w *Window) Swap() {
    sdl.GL_SwapWindow(w.Window)
}

// Clears the openGL window.
// Call this before everything is rendered.
func (w *Window) Clear() {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// Destroys the Bronson window.
func (w *Window) Destroy() {
    sdl.GL_DeleteContext(w.Context)
    w.Window.Destroy()
    sdl.Quit()
}

func initSDL() {
    if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
        panic(err)
    }
}

func initGL() {
    if err := gl.Init(); err != nil {
        panic(err)
    }
}

func (w *Window) createWindow() {
    var err error
    w.Window, err = sdl.CreateWindow(w.Title,
                                     sdl.WINDOWPOS_UNDEFINED,
                                     sdl.WINDOWPOS_UNDEFINED,
                                     int(w.Width), int(w.Height), sdl.WINDOW_OPENGL)
    if err != nil {
        panic(err)
    }
}

func (w *Window) createContext() {
    var err error
    w.Context, err = sdl.GL_CreateContext(w.Window)
    if err != nil {
        panic(err)
    }
}

func (w *Window) setVideoOptions() {
    gl.ClearColor(0.4,0.0,0.3,1.0)

    projectionMode()
    setViewport(w.Width, w.Height)

    //Required for per-pixel placing.
    gl.Translatef(0.375, 0.375, 0.0)

    modelViewMode()
    setModelViewOptions()
}

func projectionMode() {
    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
}

func modelViewMode() {
    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()
}

func setViewport(width int, height int) {
    gl.Viewport(0, 0, int32(width), int32(height))

    gl.Ortho(-float64(width)/2, float64(width)/2,
             -float64(height)/2, float64(height)/2, -1.0, 1.0)
}

func setModelViewOptions() {
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