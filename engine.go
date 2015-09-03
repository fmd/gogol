package bronson

import (
)

type Bronson struct {
    Window *Window
    EventBuffer EventBuffer
}

func New(winOpts WindowOpts) *Bronson {
    b := &Bronson{
        Window: NewWindow(winOpts),
        EventBuffer: EventBuffer{},
    }

    BuildAFuckingSquare()

    return b
}

func (e *Bronson) ProcessOneFrame() {
    e.Window.Swap()
    ProcessOneFrameOfInput(&e.EventBuffer)
    e.Window.Clear()
    RenderTheFuckingSquare()
}

func (e *Bronson) ReceiveEvents() EventBuffer {
    buffer := e.EventBuffer
    e.EventBuffer = EventBuffer{}
    return buffer
}

func (e *Bronson) Cleanup() {
    e.Window.Destroy()
}