package gogol

import (
    "github.com/go-gl/gl/v2.1/gl"
)

type Transform struct {
    position Vector
    scale    Vector
    rotation float32
    parentTransform *Transform
    childTransforms []*Transform

    Matrix Matrix
    NeedsUpdate bool
}

var rootTransform *Transform

func init() {
    rootTransform = &Transform{
        parentTransform: nil,
        Matrix: NewMatrix(),
    }
}

func NewTransform(parent *Transform) *Transform {
    t := &Transform{
        parentTransform: parent,
        Matrix: NewMatrix(),
    }

    if parent == nil {
        t.parentTransform = rootTransform
    }

    return t
}

func (t *Transform) AddChild(c *Transform) {
    c.parentTransform = t
    t.childTransforms = append(t.childTransforms, c)
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
    for _, c := range t.childTransforms {
        c.SetNeedsUpdate()
    }
}

func (t *Transform) Update() {
    var m *float32
    if t.parentTransform != nil {
        m = &(t.parentTransform.Matrix[0])
    } else {
        m = &(t.Matrix[0])
    }

    gl.LoadMatrixf(m)
    gl.Translatef(t.position.X, t.position.Y, 0)
    gl.Rotatef(t.rotation, 0, 0, -1)
    gl.GetFloatv(gl.MODELVIEW_MATRIX, &(t.Matrix[0]))
    t.NeedsUpdate = false

    for _, c := range t.childTransforms {
        c.Update()
    }
}