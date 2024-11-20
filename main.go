package main

import (
	"fmt"

	"util/util" // Adjust the import path as per your module name
)

func main() {
	fmt.Println("Demo: IsPrime Function")
	numbers := []int{2, 4, 7, 9, 11}
	for _, num := range numbers {
		fmt.Printf("Is %d prime? %v\n", num, util.IsPrime(num))
	}

	fmt.Println("\nDemo: ReverseString Function")
	str := "hello world"
	fmt.Printf("Original: %s, Reversed: %s\n", str, util.ReverseString(str))

	fmt.Println("\nDemo: SumOfSlice Function")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v, Sum: %d\n", slice, util.SumOfSlice(slice))

	fmt.Println("\nDemo: Factorial Function")
	for i := 0; i <= 5; i++ {
		fmt.Printf("Factorial of %d: %d\n", i, util.Factorial(i))
	}
}
