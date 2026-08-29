package main

import "fmt"

type FastService struct {
    state int
}

func (s *FastService) load_resolver(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*55) % 997
    }
    return total
}

func main() {
    obj := &FastService{state: 55}
    fmt.Println(obj.load_resolver(55))
}
