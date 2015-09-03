package bronson

import (
    "github.com/go-gl/gl/v2.1/gl"
    //"fmt"
)

func BuildAFuckingSquare() {
    //We're well past triangles now, baby. Back to OpenGL 2.1 for that feelgood retro aesthetic.
    // var t [4096]float32
    //var buffer uint32
    //gl.GenBuffers(1, &buffer)
}

func RenderTheFuckingSquare() {
    gl.Begin(gl.QUADS)
    gl.Color3f(1.0, 0.0, 0.0)
    gl.Vertex2f(-50.0, -50.0)
    gl.Color3f(0.0, 1.0, 0.0)
    gl.Vertex2f(-50.0, 50.0)
    gl.Color3f(0.0, 0.0, 1.0)
    gl.Vertex2f(50.0, 50.0)
    gl.Color3f(1.0, 1.0, 1.0)
    gl.Vertex2f(50.0, -50.0)
    gl.End()

    gl.Begin(gl.QUADS)
    gl.Color3f(1.0, 0.0, 0.0)
    gl.Vertex2f(51.0, -50.0)
    gl.Color3f(0.0, 1.0, 0.0)
    gl.Vertex2f(51.0, 50.0)
    gl.Color3f(0.0, 0.0, 1.0)
    gl.Vertex2f(151.0, 50.0)
    gl.Color3f(1.0, 1.0, 1.0)
    gl.Vertex2f(151.0, -50.0)
    gl.End()
}

/*
const NUM_VERTICES_IN_SHAPE uint32 = 4
const FLOAT_SIZE uint32 = 4

type VertexBuffer struct {
    Index uint32
    Object uint32
}

var VertexBuffers []*VertexBuffer
var AllGeometry []*Geometry

func NewVertexBuffer() *VertexBuffer {
    o := &VertexBuffer{
        Index: 0,
    }

    var t [4096]float32

    gl.GenBuffersARB(1, &o.Object)
    gl.BindBufferARB(gl.ARRAY_BUFFER_ARB, o.Object)
    gl.BufferDataARB(gl.ARRAY_BUFFER_ARB, int(uint32(len(t)) * FLOAT_SIZE), unsafe.Pointer(&t), gl.STATIC_DRAW)

    VertexBuffers = append(VertexBuffers, o)
    return o
}

func (o *VertexBuffer) Destroy() {
    gl.DeleteBuffers(1,&o.Object)
}

type GeometryOpts struct {
    Width float32
    Height float32
}

type Geometry struct {
    Width float32
    Height float32
    VertexBuffer *VertexBuffer
    Index uint32
}

func NewGeometry(opts GeometryOpts) *Geometry {
    var vbo *VertexBuffer
    if len(VertexBuffers) == 0 {
        vbo = NewVertexBuffer()
    }

    vbo = VertexBuffers[len(VertexBuffers)-1]

    if vbo.Index + NUM_VERTICES_IN_SHAPE > 4096 {
        vbo = NewVertexBuffer()
        VertexBuffers = append(VertexBuffers, vbo)
    }

    g := &Geometry{
        Width: opts.Width,
        Height: opts.Height,
        VertexBuffer: vbo,
        Index: vbo.Index,
    }

    addGeometryToBuffer(g, vbo)

    vbo.Index += 4
    AllGeometry = append(AllGeometry, g)

    return g
}

func addGeometryToBuffer(g *Geometry, vbo *VertexBuffer) {
    v_lx := 0.0 - g.Width/2
    v_by := 0.0 - g.Height/2

    v_rx := v_lx + g.Width
    v_ty := v_by + g.Height

    v := []float32{v_lx,v_by,v_rx,v_by,v_rx,v_ty,v_lx,v_ty}

    gl.BindBufferARB(gl.ARRAY_BUFFER_ARB, vbo.Object)
    gl.BufferSubDataARB(gl.ARRAY_BUFFER_ARB, 0, int(FLOAT_SIZE * 4), unsafe.Pointer(&v))
}

func RenderAllGeometry() {
    var vbo uint32 = 1

    gl.BindBufferARB(gl.ARRAY_BUFFER_ARB, vbo)
    gl.VertexPointer(2, gl.FLOAT, 0, nil)

    gl.EnableClientState(gl.VERTEX_ARRAY)

    for _, g := range AllGeometry {
        render(g, &vbo)
    }

    gl.DisableClientState(gl.VERTEX_ARRAY)
}

func render(g *Geometry, vbo *uint32) {
    if *vbo != g.VertexBuffer.Object {
        vbo = &g.VertexBuffer.Object

        gl.BindBufferARB(gl.ARRAY_BUFFER_ARB, *vbo)
        gl.VertexPointer(2, gl.FLOAT, 0, nil)
    }

    gl.DrawArrays(uint32(gl.QUADS), int32(g.Index), int32(NUM_VERTICES_IN_SHAPE))
}*/