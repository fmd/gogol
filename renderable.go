package gogol

import (
    "container/list"
    "github.com/go-gl/gl/v2.1/gl"
)

type Renderable struct {
    VboPosition VboPosition
    Element *list.Element
    Layer *Layers
    Length int
}

func NewRenderable(verts []float32) *Renderable {
    r := &Renderable{
        VboPosition: G.Renderer.NewVboPosition(verts),
        Length: len(verts),
    }

    gl.BindBuffer(gl.ARRAY_BUFFER, r.VboPosition.Vbo.Id)
    gl.BufferSubData(gl.ARRAY_BUFFER, int(r.VboPosition.Index*4), len(verts)*4, gl.Ptr(verts))
    return r
}

func (r *Renderable) MoveAfter(c Component) {
    if c.GetRenderable().Layer != r.Layer {
        return
    }

    r.Layer.MoveAfter(r.Element, c.GetRenderable().Element)
}

func (r *Renderable) MoveBefore(c Component) {
    if c.GetRenderable().Layer != r.Layer {
        return
    }

    r.Layer.MoveBefore(r.Element, c.GetRenderable().Element)
}

func (r *Renderable) MoveToFront() {
    r.Layer.MoveToFront(r.Element)
}

func (r *Renderable) MoveToBack() {
    r.Layer.MoveToBack(r.Element)
}