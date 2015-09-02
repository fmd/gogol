package engine

import (
    "github.com/veandco/go-sdl2/sdl"
)

var event sdl.Event
type EventBuffer []sdl.Event

func ProcessOneFrameOfInput(buffer *EventBuffer) {
    for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
        *buffer = append(*buffer, event)
    }
}