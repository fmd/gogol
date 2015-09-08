package gogol

// The global Gogol object.
var G *Gogol

// The type structure of the global Gogol object.
// There should only ever be one of these.
type Gogol struct {
    Quit bool
    Window *Window
    Input *Input
    Renderer *Renderer
}

// Initialises Gogol.
// This function is not called 'New' to signify it should only be called once.
func Init(opts WindowOpts) *Gogol {
    G = &Gogol{
        Window: NewWindow(opts),
        Input: NewInput(),
        Renderer: NewRenderer(),
    }

    return G
}

// A global function that returns whether Gogol should quit.
func ShouldQuit() bool {
    return G.Quit
}

// Processes one frame.
// This should be called on each pass of the game's loop.
func ProcessOneFrame() {
    G.Input.Handle()
    G.Window.Swap()
    G.Window.Clear()
    G.Renderer.Render()
}

// Quits Gogol. Defer this to the end of your app's main function.
// Destroys the window, cleans up the openGL context.
func Cleanup() {
    G.Window.Destroy()
}