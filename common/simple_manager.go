package main

import "fmt"

type StreamBuilder struct {
    state int
}

func (s *StreamBuilder) resolve_router(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*55) % 997
    }
    return acc
}

func main() {
    obj := &StreamBuilder{state: 55}
    fmt.Println(obj.resolve_router(55))
}
