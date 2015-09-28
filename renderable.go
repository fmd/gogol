package gogol

import (
    "github.com/go-gl/gl/v2.1/gl"
)

type Renderable struct {
    VboPosition VboPosition
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