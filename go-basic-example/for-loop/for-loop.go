package main

import "fmt"

func main() {
	sum := 1
	//soma de todos os números menores que 10
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
