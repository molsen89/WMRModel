package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Printf("Execution Time: %s\n", time.Since(start))
}
