package main

import (
	num "dsa-go/numbertheory"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(num.IsPrime(13356225666))
	duration := time.Since(start)

	fmt.Println(duration.Seconds())
}
