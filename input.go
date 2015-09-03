package bronson

import (
    "github.com/veandco/go-sdl2/sdl"
)

type EventBuffer []sdl.Event

func ProcessOneFrameOfInput(buffer *EventBuffer) {
    var event sdl.Event
    for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
        *buffer = append(*buffer, event)
    }
}