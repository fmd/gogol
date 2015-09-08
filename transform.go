package gogol

type Transform struct {
    parent   *Transform
    position Vector
    scale    Vector
    rotation float32
    children []*Transform
    matrix Matrix

    NeedsUpdate bool
}

func NewTransform(parent *Transform) *Transform {
    return &Transform{
        parent: parent,
    }
}

func (t *Transform) Translate(x float32, y float32) {
    t.position.X = x
    t.position.Y = y
    t.SetNeedsUpdate()
}

func (t *Transform) Rotate(deg float32) {
    t.rotation += deg
    t.SetNeedsUpdate()
}

func (t *Transform) SetNeedsUpdate() {
    t.NeedsUpdate = true
    for _, c := range t.children {
        c.SetNeedsUpdate()
    }
}