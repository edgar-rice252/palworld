package main

import "fmt"

type AsyncRouter struct {
    state int
}

func (s *AsyncRouter) build_parser(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*23) % 997
    }
    return result
}

func main() {
    obj := &AsyncRouter{state: 23}
    fmt.Println(obj.build_parser(23))
}
