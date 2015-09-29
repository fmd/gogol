package gogol

import (
    "github.com/go-gl/gl/v2.1/gl"
)

const (
    ValsInVertex = 2
    VboLength = 4096
)

type Renderer struct {
    Vbos []*Vbo
    VboPosition VboPosition
    Layer *Layer
}

func NewRenderer() *Renderer {
    vbo := NewVbo(VboLength)
    r := &Renderer{
        Vbos: []*Vbo{vbo},
        VboPosition: VboPosition{
            Vbo: vbo,
            Index: 0,
        },
        Layer: NewLayer("base"),
    }
    return r
}

func (r *Renderer) Render() {
    gl.EnableClientState(gl.VERTEX_ARRAY)

    rVbo := r.Vbos[0]
    gl.BindBuffer(gl.ARRAY_BUFFER, rVbo.Id)

    for _, c := range r.Layer.Flatten() {
        renderable := c.GetRenderable()
        transform := c.GetTransform()

        if renderable == nil || transform == nil || !c.Visible() {
            continue
        }

        if rVbo != renderable.VboPosition.Vbo {
            rVbo = renderable.VboPosition.Vbo
            gl.BindBuffer(gl.ARRAY_BUFFER, rVbo.Id)
        }

        gl.VertexPointer(ValsInVertex, gl.FLOAT, 0, nil)

        if transform.NeedsUpdate {
            transform.Update()
        }

        gl.LoadMatrixf(&(transform.Matrix[0]))
        gl.DrawArrays(gl.QUADS, int32(renderable.VboPosition.Index / ValsInVertex),
                                int32(renderable.Length / ValsInVertex))
    }

    gl.DisableClientState(gl.VERTEX_ARRAY);
}

func (r *Renderer) NewVboPosition(verts []float32) VboPosition {
    pos := r.copyVboPosition()

    if len(verts) + int(r.VboPosition.Index) > VboLength {
        vbo := NewVbo(VboLength)
        r.Vbos = append(r.Vbos, vbo)
        r.VboPosition = VboPosition{
            Vbo: vbo,
            Index: 0,
        }
    } else {
        r.VboPosition.Index += uint32(len(verts))
    }

    return pos
}

func (r *Renderer) copyVboPosition() VboPosition{
    return VboPosition{
        Vbo: r.VboPosition.Vbo,
        Index: r.VboPosition.Index,
    }
}