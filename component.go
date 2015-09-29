package gogol

//TODO: Maybe rethink this?
type Component interface {
    GetTransform() *Transform
    GetRenderable() *Renderable
    Visible() bool

    SetParent(Component)
    AddChild(Component)
}