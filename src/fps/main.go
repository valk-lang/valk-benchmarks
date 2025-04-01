package main

import (
    "fmt"
    "time"
    "runtime"
)

type Node struct {
    next *Node
	data int
}

func createObjects(amount int) *Node {
    first := &Node{ data: 0 }
    r := first
    for i := 0; i < amount; i++ {
		n := &Node{ data: i }
		r.next = n
		r = n
    }
	return first
}

func main() {
    fmt.Println("Running game loop...")
    totalFrames := 500

	var liveObjects []*Node
	liveObjects = append(liveObjects, createObjects(10 * 1000 * 1000))

    start := time.Now()
    for frame := 0; frame < totalFrames; frame++ {
		runtime.GC()
        liveObjects = append(liveObjects, createObjects(1))
    }

    elapsed := time.Since(start)

    fmt.Printf("Kept %d node chains alive\n", len(liveObjects))
    fmt.Printf("Finished %d frames in %s\n", totalFrames, elapsed)
    fmt.Printf("FPS: %d\n", totalFrames * 1000 / int(elapsed / time.Millisecond + 1))
}
