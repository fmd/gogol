package gogol

import (
    "github.com/go-gl/gl/v2.1/gl"
)

const (
    FloatSize = 4
)

type VboPosition struct {
    Vbo *Vbo
    Index uint32
}

type Vbo struct {
    Vertices []float32
    Id       uint32
}

func NewVbo(length int) *Vbo {
    v := &Vbo{
        Vertices: make([]float32, length),
    }

    gl.GenBuffers(1, &v.Id)
    gl.BindBuffer(gl.ARRAY_BUFFER, v.Id)
    gl.BufferData(gl.ARRAY_BUFFER, length * FloatSize, gl.Ptr(v.Vertices), gl.STATIC_DRAW)
    return v
}