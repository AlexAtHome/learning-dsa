package main

import "alexalex.dev/dsa/stack"

func findSum() (sum int) {
	sum = 0
	for i := 1; i < 100000000000; i++ {
		sum += i
	}
	return sum
}

func findSum2(input int) int {
	return input * (input + 1) / 2
}

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func reverse[K any](arr []K, start int, end int) []K {
	if start >= end {
		return arr
	}

	temp := arr[start]
	arr[start] = arr[end]
	arr[end] = temp

	return reverse(arr, start+1, end-1)
}

/** Time complexity: O(1) */
func rotateRight[K any](arr []K, steps int) []K {
	return append(arr[len(arr)-steps:], arr[:len(arr)-steps]...)
}

/** Time complexity: O(1) */
func rotateLeft[K any](arr []K, steps int) []K {
	return append(arr[steps:], arr[:steps]...)
}

func main() {
	// array := []int{1, 2, 3, 4, 5, 6}

	// fmt.Println(array)
	// fmt.Println("rotateRight", rotateRight(array, 2))
	// fmt.Println("rotateLeft", rotateLeft(array, 2))
	// fmt.Println("reverse", reverse(array, 0, 5))
	//
	// fmt.Println("")
	//
	// array2 := []string{"foo", "bar", "baz", "bic", "far"}
	// fmt.Println(array2)
	// fmt.Println("rotateRight", rotateRight(array2, 2))
	// fmt.Println("rotateLeft", rotateLeft(array2, 2))
	// fmt.Println("reverse", reverse(array2, 0, 4))
	stack.StackTest()
}
