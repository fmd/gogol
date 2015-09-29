package gogol

type BaseComponent struct {
    visible bool

    *Transform
    *Renderable
    Parent Component
    Children []Component
}

func NewBaseComponent(verts []float32, parent *Transform) *BaseComponent {
    c := &BaseComponent{
        Transform: NewTransform(parent),
        Renderable: NewRenderable(verts),
        visible: true,
    }

    c.AddToLayer("default")
    return c
}

func (b *BaseComponent) AddToLayer(name string) {
    layer := G.Renderer.RenderList.GetLayer(name)

    if b.Renderable.Layer != nil {
        b.Renderable.Layer.Remove(b.Renderable.Element)
    }

    b.Renderable.Element = layer.PushBack(b)
    b.Renderable.Layer = layer
}

func (b *BaseComponent) GetTransform() *Transform {
    return b.Transform
}

func (b *BaseComponent) GetRenderable() *Renderable {
    return b.Renderable
}

func (b *BaseComponent) Visible() bool {
    return b.visible
}

func (b *BaseComponent) Show() {
    b.visible = true
    for _, c := range b.Children {
        c.Show()
    }
}

func (b *BaseComponent) Hide() {
    b.visible = true
    for _, c := range b.Children {
        c.Hide()
    }
}

func (b *BaseComponent) SetParent(c Component) {
    b.Parent = c
}

func (b *BaseComponent) AddChild(c Component) {
    c.SetParent(b)
    b.Children = append(b.Children, c)
    b.Transform.AddChild(c.GetTransform())
}