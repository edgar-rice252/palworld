package main

import "fmt"

type BatchSession struct {
    state int
}

func (s *BatchSession) render_builder(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*26) % 997
    }
    return result
}

func main() {
    obj := &BatchSession{state: 26}
    fmt.Println(obj.render_builder(26))
}
