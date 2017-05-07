package main

import (
	"./directedgraphs"
	"fmt"
)

func main() {
	input := []string{"AB5", "BC4", "CD8", "DC8", "DE6", "AD5", "CE2", "EB3", "AE7"}
	fmt.Printf("Output#1:%s \n", directedgraphs.Problem1(input))
	fmt.Printf("Output#2:%s \n", directedgraphs.Problem2(input))
	fmt.Printf("Output#3:%s \n", directedgraphs.Problem3(input))
	fmt.Printf("Output#4:%s \n", directedgraphs.Problem4(input))
	fmt.Printf("Output#5:%s \n", directedgraphs.Problem5(input))
}
