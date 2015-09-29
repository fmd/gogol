package gogol

import (
    "container/list"
)

var LayerMap map[string]*Layers

func init() {
    LayerMap = make(map[string]*Layers)
}

func GetLayers(name string) *Layers {
    return LayerMap[name]
}

type Layers struct {
    *list.List
    visible bool
}

func NewLayers() *Layers {
    l := &Layers{
        List: list.New(),
    }

    return l
}

func (l *Layers) Visible() bool {
    return l.visible
}

func (l *Layers) Show() {
    l.visible = true
}

func (l *Layers) Hide() {
    l.visible = false
}

func (l *Layers) AppendLayers(name string) {
    LayerMap[name] = &Layers{
        List: list.New(),
        visible: true,
    }

    l.PushBack(LayerMap[name])
}

func (l *Layers) PrependLayers(name string) {
    LayerMap[name] = &Layers{
        List: list.New(),
        visible: true,
    }

    l.PushBack(LayerMap[name])
}

func (l *Layers) Flatten() []Component {
    var components []Component
    for el := l.Front(); el != nil; el = el.Next() {
        layer, ok := el.Value.(*Layers)
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