package main

import (
	num "dsa-go/numbertheory"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	num.Seive(2023)
	duration := time.Since(start)

	fmt.Println("Execution time: ", duration.Milliseconds())
}
