package engine

type Engine struct {
    Window *Window
    EventBuffer EventBuffer
}

func NewEngine(winOpts WindowOpts) *Engine {
    e := &Engine{
        Window: NewWindow(winOpts),
        EventBuffer: EventBuffer{},
    }

    return e
}

func (e *Engine) ProcessOneFrame() {
    e.Window.Swap()
    ProcessOneFrameOfInput(&e.EventBuffer)
    e.Window.Clear()
    RenderAllRenderables()
}

func (e *Engine) ReceiveEvents() EventBuffer {
    buffer := e.EventBuffer
    e.EventBuffer = EventBuffer{}
    return buffer
}