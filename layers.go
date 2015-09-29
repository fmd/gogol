package gogol

import (
    "container/list"
)

var LayerMap map[string]*Layer

func init() {
    LayerMap = make(map[string]*Layer)
}

func GetLayer(name string) *Layer {
    return LayerMap[name]
}

type Layer struct {
    *list.List
    visible bool
}

func NewLayer(name string) *Layer {
    l := &Layer{
        List: list.New(),
        visible: true,
    }

    LayerMap[name] = l
    return l
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

func (l *Layer) AppendLayer(name string) {
    LayerMap[name] = &Layer{
        List: list.New(),
        visible: true,
    }

    l.PushBack(LayerMap[name])
}

func (l *Layer) PrependLayer(name string) {
    LayerMap[name] = &Layer{
        List: list.New(),
        visible: true,
    }

    l.PushBack(LayerMap[name])
}

func (l *Layer) Flatten() []Component {
    var components []Component
    for el := l.Front(); el != nil; el = el.Next() {
        layer, ok := el.Value.(*Layer)
        if !ok {
            if element, ok := el.Value.(Component); ok {
                components = append(components, element)
            }
            continue
        }

        if !layer.Visible() {
            continue
        }

        components = append(components, layer.Flatten()...)
    }

    return components
}