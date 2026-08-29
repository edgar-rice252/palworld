package main

import "fmt"

type FastHandler struct {
    state int
}

func (s *FastHandler) handle_adapter(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*40) % 997
    }
    return value
}

func main() {
    obj := &FastHandler{state: 40}
    fmt.Println(obj.handle_adapter(40))
}
