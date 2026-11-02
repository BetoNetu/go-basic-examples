package main

import (
	"time"
)

// Goo structure for the final loop of the daemon
// just switch the condition of the if statement
func main() {
	for count := 10; count > 0; count -= 1 {
		time.Sleep(1 * time.Second)
		println(count)
		if count == 1 {
			count += 10
		}
	}

}
