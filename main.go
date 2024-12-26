package main

import (
	"fmt"
	"github.com/ethanhorrigan/kubeskope/cmd"
	"time"
)

func main() {
	start := time.Now()
	cmd.Execute()
	fmt.Printf("Execution time: %v\n", time.Since(start))
}
