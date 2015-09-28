package gogol

type BaseComponent struct {
    *Transform
    *Renderable
    parent Component
    children []Component
}

func (b *BaseComponent) GetTransform() *Transform {
    return b.Transform
}

func (b *BaseComponent) GetRenderable() *Renderable {
    return b.Renderable
}

//TODO: Change, add scene add, scene remove, etc
func (b *BaseComponent) ShouldRender() bool {
    return true
}

//TODO: Maybe rethink this?
func (b *BaseComponent) SetParent(c Component) {
    b.parent = c
}

//TODO: Maybe rethink this?
func (b *BaseComponent) AddChild(c Component) {
    c.SetParent(b)
    b.children = append(b.children, c)
    b.Transform.AddChild(c.GetTransform())
}

func NewBaseComponent(verts []float32, parent *Transform) *BaseComponent {
    c := &BaseComponent{
        Transform: NewTransform(parent),
        Renderable: NewRenderable(verts),
    }

    //TODO: This should be kept track of and added and removed with the scene
    G.Renderer.Components = append(G.Renderer.Components, c)
    return c
}