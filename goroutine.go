package main
import (
	"fmt"
	"time"
)

func squareIt(x int) {
	fmt.Println(x * x)
}

func main() {
	go squareIt(2)	// execute 'squareIt' in the background (another thread)
	time.Sleep(1 * time.Millisecond)	// wait
}