package gogol

import (
    "github.com/veandco/go-sdl2/sdl"
)

type KeyEventHandler []func(KeyCode)
type InputEventHandlers map[KeyCode][]func(KeyCode)

type Input struct {
    KeyDownHandlers InputEventHandlers
    KeyUpHandlers InputEventHandlers
    MouseDownHandlers InputEventHandlers
    MouseUpHandlers InputEventHandlers
    KeyStates []uint8
    oldKeyStates []uint8
}

func NewInput() *Input {
    i := &Input{
        KeyDownHandlers: make(InputEventHandlers),
        KeyUpHandlers: make(InputEventHandlers),
        MouseDownHandlers: make(InputEventHandlers),
        MouseUpHandlers: make(InputEventHandlers),
        KeyStates: sdl.GetKeyboardState(),
    }
    i.oldKeyStates = make([]uint8, len(i.KeyStates))

    return i
}

func OnMouseDown(code KeyCode, fn func(KeyCode)) {
    G.Input.MouseDownHandlers[code] = append(G.Input.MouseDownHandlers[code], fn)
}

func OnKeyDown(code KeyCode, fn func(KeyCode)) {
    G.Input.KeyDownHandlers[code] = append(G.Input.KeyDownHandlers[code], fn)
}

func OnKeyUp(code KeyCode, fn func(KeyCode)) {
    G.Input.KeyUpHandlers[code] = append(G.Input.KeyUpHandlers[code], fn)
}

func (i *Input) Handle() {
    var event sdl.Event

    copy(i.oldKeyStates, i.KeyStates)
    for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
        switch e := event.(type) {
            case *sdl.QuitEvent:
                G.Quit = true
            case *sdl.MouseButtonEvent:
                i.handleMouseButtonEvent(e)
            case *sdl.KeyDownEvent:
                i.handleKeyDownEvent(e)
            case *sdl.KeyUpEvent:
                i.handleKeyUpEvent(e)
        }
    }
}

func (i *Input) handleMouseButtonEvent(e *sdl.MouseButtonEvent) {
    t := e.Type
    code := KeyCode(e.Button)
    if t == sdl.MOUSEBUTTONDOWN {
        i.MouseDownHandlers.trigger(code)
    } else if t == sdl.MOUSEBUTTONUP {
        i.MouseUpHandlers.trigger(code)
    }
}

func (i *Input) handleKeyDownEvent(e *sdl.KeyDownEvent) {
    code := sdl.GetScancodeFromKey(e.Keysym.Sym)
    if i.oldKeyStates[code] == 0 {
        i.KeyDownHandlers.trigger(KeyCode(code))
    }
}

func (i *Input) handleKeyUpEvent(e *sdl.KeyUpEvent) {
    code := sdl.GetScancodeFromKey(e.Keysym.Sym)
    if i.oldKeyStates[code] == 1 {
        i.KeyUpHandlers.trigger(KeyCode(code))
    }
}

func (m InputEventHandlers) trigger(code KeyCode) {
    h := m[code]
    for _, fn := range h {
        fn(code)
    }
}