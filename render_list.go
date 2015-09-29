package gogol

import (
    "container/list"
)

type Layer struct {
    *list.List
    visible bool
}

func (l *Layer) Visible() bool {
    return l.visible
}

func (l *Layer) Show() {
    l.visible = true
}

func (l *Layer) Hide() {
    l.visible = false
}

type RenderList struct {
    layers map[string]*Layer
    list *list.List
}

func NewRenderList() *RenderList {
    l := &RenderList{
        layers: make(map[string]*Layer),
        list: list.New(),
    }

    l.AppendLayer("default")
    return l
}

func (l *RenderList) AppendLayer(name string) {
    l.layers[name] = &Layer{
        List: list.New(),
        visible: true,
    }

    l.list.PushBack(l.layers[name])
}

func (l *RenderList) PrependLayer(name string) {
    l.layers[name] = &Layer{
        List: list.New(),
        visible: true,
    }

    l.list.PushBack(l.layers[name])
}

func (l *RenderList) GetLayer(name string) *Layer {
    return l.layers[name]
}

func (l *RenderList) Flatten() []Component {
    var components []Component
    for la := l.list.Front(); la != nil; la = la.Next() {
        layer := la.Value.(*Layer)
        if !layer.Visible() {
            continue
        }

        for el := layer.Front(); el != nil; el = el.Next() {
            components = append(components, el.Value.(Component))
        }
    }

    return components
}