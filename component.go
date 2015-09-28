package gogol

//TODO: Maybe rethink this?
type Component interface {
    GetTransform() *Transform
    GetRenderable() *Renderable
    ShouldRender() bool

    SetParent(Component)
    AddChild(Component)
}