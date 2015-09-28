package gogol

type Quad struct {
    *BaseComponent
}

type QuadOpts struct {
    Width float32
    Height float32
}

func NewQuad(opts QuadOpts) *Quad {
    halfWidth := opts.Width / 2
    halfHeight := opts.Height / 2

    verts := []float32{
        -halfWidth, -halfHeight,
        -halfWidth, halfHeight,
        halfWidth, halfHeight,
        halfWidth, -halfHeight,
    }

    return &Quad{
        BaseComponent: NewBaseComponent(verts, nil),
    }
}