package gogol

type Vector struct {
    X float32
    Y float32
}

var (
    Zero     = Vector{0, 0}
    Up       = Vector{0, 1}
    Down     = Vector{0, -1}
    Left     = Vector{-1, 0}
    Right    = Vector{1, 0}
)

