package main

import "fmt"

type AsyncProcessor struct {
    state int
}

func (s *AsyncProcessor) dispatch_adapter(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*66) % 997
    }
    return count
}

func main() {
    obj := &AsyncProcessor{state: 66}
    fmt.Println(obj.dispatch_adapter(66))
}
