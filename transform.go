package gogol

type Transform struct {
    parent   *Transform
    position Vector
    rotation Vector
    scale    Vector

    children []*Transform
    matrix Matrix
}

func NewTransform(parent *Transform) *Transform {
    return &Transform{
        parent: parent,
    }
}