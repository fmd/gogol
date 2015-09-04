package bronson

import (
    "github.com/veandco/go-sdl2/sdl"
)

type KeyEventHandler []func(KeyCode)
type KeyEventHandlerMap map[KeyCode][]func(KeyCode)

type Input struct {
    KeyDownHandlers KeyEventHandlerMap
    KeyUpHandlers KeyEventHandlerMap
    KeyStates []uint8
    oldKeyStates []uint8
}

func NewInput() *Input {
    i := &Input{
        KeyDownHandlers: make(KeyEventHandlerMap),
        KeyUpHandlers: make(KeyEventHandlerMap),
        KeyStates: sdl.GetKeyboardState(),
    }
    i.oldKeyStates = make([]uint8, len(i.KeyStates))

    return i
}


func (i *Input) Handle() {
    var event sdl.Event

    copy(i.oldKeyStates, i.KeyStates)
    for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
        switch t := event.(type) {
            case *sdl.QuitEvent:
                B.Quit = true
            case *sdl.KeyDownEvent:
                code := sdl.GetScancodeFromKey(t.Keysym.Sym)
                if i.oldKeyStates[code] == 0 {
                    i.KeyDownHandlers.trigger(KeyCode(code))
                }
            case *sdl.KeyUpEvent:
                code := sdl.GetScancodeFromKey(t.Keysym.Sym)
                if i.oldKeyStates[code] == 1 {
                    i.KeyUpHandlers.trigger(KeyCode(code))
                }
        }
    }
}

func OnKeyDown(code KeyCode, fn func(KeyCode)) {
    B.Input.KeyDownHandlers[code] = append(B.Input.KeyDownHandlers[code], fn)
}

func OnKeyUp(code KeyCode, fn func(KeyCode)) {
    B.Input.KeyUpHandlers[code] = append(B.Input.KeyUpHandlers[code], fn)
}

func (m KeyEventHandlerMap) trigger(code KeyCode) {
    h := m[code]
    for _, fn := range h {
        fn(code)
    }
}