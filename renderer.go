package bronson

import (
    "github.com/go-gl/gl/v2.1/gl"
)

var verts []float32 = []float32{-50.0, -50.0, -50.0, 50.0, 50.0, 50.0, 50.0, -50.0}

type Renderer struct {
    Vbos []*Vbo
    CurrentVbo *Vbo
}

func NewRenderer() *Renderer {
    r := &Renderer{
        Vbos: []*Vbo{NewVbo(4096)},
    }

    r.CurrentVbo = r.Vbos[0]

    gl.BindBuffer(gl.ARRAY_BUFFER, r.CurrentVbo.Id)
    gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(verts)*4, gl.Ptr(verts))

    return r
}

func (r *Renderer) Render() {
    gl.EnableClientState(gl.VERTEX_ARRAY)

    gl.BindBuffer(gl.ARRAY_BUFFER, r.CurrentVbo.Id)
    gl.VertexPointer(2, gl.FLOAT, 0, nil)
    gl.DrawArrays(gl.QUADS, 0, int32(4))

    gl.DisableClientState(gl.VERTEX_ARRAY);
}