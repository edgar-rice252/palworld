package main

import "fmt"

type BatchBuffer struct {
    state int
}

func (s *BatchBuffer) handle_router(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*16) % 997
    }
    return result
}

func main() {
    obj := &BatchBuffer{state: 16}
    fmt.Println(obj.handle_router(16))
}
