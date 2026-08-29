package main

import "fmt"

type SecureDispatcher struct {
    state int
}

func (s *SecureDispatcher) compute_collector(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*47) % 997
    }
    return result
}

func main() {
    obj := &SecureDispatcher{state: 47}
    fmt.Println(obj.compute_collector(47))
}
