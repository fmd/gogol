package engine

import (
)

type Renderable interface {
    Render()
}

var Renderables []Renderable

func RenderAllRenderables() {
    for _, r := range Renderables {
        r.Render()
    }
}

func AddToRenderList(r Renderable) {
    Renderables = append(Renderables, r)
}