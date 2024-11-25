package main

import (
	"fmt"
	"project_1/DataStructures" 
)

func main() {
	// Priority Queue for integers
	intPQ := DataStructures.NewPriorityQueue(func(a, b interface{}) bool {
		return a.(int) < b.(int) // Compare integers
	})

	intPQ.Push(5)
	intPQ.Push(3)
	intPQ.Push(8)
	intPQ.Push(1)

	if value, err := intPQ.Peek(); err == nil {
		fmt.Println(value) // Should print 1
	} else {
		fmt.Println(err)
	}

	if value, err := intPQ.Pop(); err == nil {
		fmt.Println(value) // Should print 1
	} else {
		fmt.Println(err)
	}
	if value, err := intPQ.Pop(); err == nil {
		fmt.Println(value) // Should print 3
	} else {
		fmt.Println(err)
	}
	if value, err := intPQ.Pop(); err == nil {
		fmt.Println(value) // Should print 5
	} else {
		fmt.Println(err)
	}
	if value, err := intPQ.Pop(); err == nil {
		fmt.Println(value) // Should print 8
	} else {
		fmt.Println(err)
	}

	fmt.Println(intPQ.IsEmpty()) // Should print true

	// Priority Queue for strings
	stringPQ := DataStructures.NewPriorityQueue(func(a, b interface{}) bool {
		return a.(string) < b.(string) // Compare strings lexicographically
	})

	stringPQ.Push("banana")
	stringPQ.Push("apple")
	stringPQ.Push("cherry")
	stringPQ.Push("date")

	if value, err := stringPQ.Peek(); err == nil {
		fmt.Println(value) // Should print "apple"
	} else {
		fmt.Println(err)
	}

	for !stringPQ.IsEmpty() {
		if value, err := stringPQ.Pop(); err == nil {
			fmt.Println(value) // Print elements in order: apple, banana, cherry, date
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println(stringPQ.IsEmpty()) //
}