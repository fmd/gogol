package gogol

import (
    "github.com/go-gl/gl/v2.1/gl"
)

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

func (t *Transform) update() {
    gl.LoadMatrixf(&(t.parent.matrix[0]))
    gl.Translatef(t.position.X, t.position.Y, 0)
    gl.Rotatef(t.rotation, 0, 0, -1)
    gl.GetFloatv(gl.MODELVIEW_MATRIX, &(t.parent.matrix[0]))

    for _, c := range t.children {
        c.update()
    }
}