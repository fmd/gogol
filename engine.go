package bronson

type Bronson struct {
    Window *Window
    EventBuffer EventBuffer
}

func New(winOpts WindowOpts) *Bronson {
    return &Bronson{
        Window: NewWindow(winOpts),
        EventBuffer: EventBuffer{},
    }
}

func (e *Bronson) ProcessOneFrame() {
    e.Window.Swap()
    ProcessOneFrameOfInput(&e.EventBuffer)
    e.Window.Clear()
    RenderAllGeometry()
}

func (e *Bronson) ReceiveEvents() EventBuffer {
    buffer := e.EventBuffer
    e.EventBuffer = EventBuffer{}
    return buffer
}