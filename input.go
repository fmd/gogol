package bronson

import (
    "github.com/veandco/go-sdl2/sdl"
)

// FIX KEY REPEAT

type KeyEventHandler []func(KeyCode)
type KeyEventHandlerMap map[KeyCode][]func(KeyCode)

type Input struct {
    KeyDownHandlers KeyEventHandlerMap
    KeyUpHandlers KeyEventHandlerMap
}

func NewInput() *Input {
    return &Input{
        KeyDownHandlers: make(KeyEventHandlerMap),
        KeyUpHandlers: make(KeyEventHandlerMap),
    }
}

func (i *Input) Handle() {
    var event sdl.Event
    for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
        switch t := event.(type) {
            case *sdl.QuitEvent:
                B.Quit = true
            case *sdl.MouseMotionEvent:
                break
                //fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
                //           t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
            case *sdl.KeyDownEvent:
                code := sdl.GetScancodeFromKey(t.Keysym.Sym)
                i.KeyDownHandlers.trigger(KeyCode(code))
            case *sdl.KeyUpEvent:
                code := sdl.GetScancodeFromKey(t.Keysym.Sym)
                i.KeyUpHandlers.trigger(KeyCode(code))
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