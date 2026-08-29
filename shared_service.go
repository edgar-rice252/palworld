package main

import "fmt"

type DynamicParser struct {
    state int
}

func (s *DynamicParser) collect_parser(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*61) % 997
    }
    return acc
}

func main() {
    obj := &DynamicParser{state: 61}
    fmt.Println(obj.collect_parser(61))
}
