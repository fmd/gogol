package bronson

// The global Bronson object.
var B *Bronson

// The type structure of the global Bronson object.
// There should only ever be one of these.
type Bronson struct {
    Quit bool
    Window *Window
    Input *Input
    Renderer *Renderer
}

// Initialises Bronson.
// This function is not called 'New' to signify it should only be called once.
func Init(opts WindowOpts) *Bronson {
    B = &Bronson{
        Window: NewWindow(opts),
        Input: NewInput(),
        Renderer: NewRenderer(),
    }

    return B
}

// A global function that returns whether Bronson should quit.
func ShouldQuit() bool {
    return B.Quit
}

// Processes one frame.
// This should be called on each pass of the game's loop.
func ProcessOneFrame() {
    B.Input.Handle()
    B.Window.Swap()
    B.Window.Clear()
    B.Renderer.RenderTheFuckingSquare()
}

// Quits Bronson. Defer this to the end of your app's main function.
// Destroys the window, cleans up the openGL context.
func Cleanup() {
    B.Window.Destroy()
}