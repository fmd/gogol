package gogol

import (
    "github.com/go-gl/gl/v2.1/gl"
)

type GameObject struct {
    Transform *Transform
    VboPosition *VboPosition
    Vertices []float32
}

type QuadOpts struct {
    Width float32
    Height float32
}

func CreateQuad(opts QuadOpts) *GameObject {
    halfWidth := opts.Width / 2
    halfHeight := opts.Height / 2

    verts := []float32{
        -halfWidth, -halfHeight,
        -halfWidth, halfHeight,
        halfWidth, halfHeight,
        halfWidth, -halfHeight,
    }

    g := &GameObject{
        Transform: NewTransform(nil),
        VboPosition: G.Renderer.GetNewVboPosition(verts),
        Vertices: verts,
    }

    gl.BindBuffer(gl.ARRAY_BUFFER, g.VboPosition.Vbo.Id)
    gl.BufferSubData(gl.ARRAY_BUFFER, int(g.VboPosition.Index*4), len(g.Vertices)*4, gl.Ptr(g.Vertices))

    G.Renderer.GameObjects = append(G.Renderer.GameObjects, g)
    return g
}