package bronson

type Transform struct {
    parent   *Transform
    position Vector
    rotation Vector
    scale    Vector

    children []*Transform
    matrix Matrix
}