package gogol

//TODO: Maybe rethink this?
type Component interface {
    GetTransform() *Transform
    GetRenderable() *Renderable

    Visible() bool
    Show()
    Hide()

    SetParent(Component)
    AddChild(Component)
}