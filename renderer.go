package gogol

import (
    "github.com/go-gl/gl/v2.1/gl"
)

const (
    VboLength = 4096
)

type Renderer struct {
    Vbos []*Vbo
    VboPosition *VboPosition
    GameObjects []*GameObject
}

func NewRenderer() *Renderer {
    vbo := NewVbo(VboLength)
    r := &Renderer{
        Vbos: []*Vbo{vbo},
        VboPosition: &VboPosition{
            Vbo: vbo,
            Index: 0,
        },
    }
    return r
}

func (r *Renderer) Render() {
    gl.EnableClientState(gl.VERTEX_ARRAY)

    rVbo := r.Vbos[0]
    gl.BindBuffer(gl.ARRAY_BUFFER, rVbo.Id)

    for _, g := range r.GameObjects {
        if rVbo != g.VboPosition.Vbo {
            rVbo = g.VboPosition.Vbo
            gl.BindBuffer(gl.ARRAY_BUFFER, rVbo.Id)
        }

        gl.VertexPointer(2, gl.FLOAT, 0, nil)

        if g.Transform.NeedsUpdate {
            g.Transform.Update()
        }

        gl.LoadMatrixf(&(g.Transform.Matrix[0]))
        gl.DrawArrays(gl.QUADS, int32(g.VboPosition.Index), int32(4))
    }

    gl.DisableClientState(gl.VERTEX_ARRAY);
}

func (r *Renderer) GetNewVboPosition(verts []float32) *VboPosition {
    pos := r.copyVboPosition()

    if len(verts) + int(r.VboPosition.Index) > VboLength {
        vbo := NewVbo(VboLength)
        r.Vbos = append(r.Vbos, vbo)
        r.VboPosition = &VboPosition{
            Vbo: vbo,
            Index: 0,
        }
    } else {
        r.VboPosition.Index += uint32(len(verts))
    }

    return pos
}

func (r *Renderer) copyVboPosition() *VboPosition{
    return &VboPosition{
        Vbo: r.VboPosition.Vbo,
        Index: r.VboPosition.Index,
    }
}