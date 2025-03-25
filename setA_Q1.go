package main

import "fmt"

func main() {
    // Example array
    numbers := []int{12, 34, 5, 87, 1, 9, 44, 23}

    // Initialize largest and smallest with the first element
    largest := numbers[0]
    smallest := numbers[0]

    // Iterate through the array to find the largest and smallest
    for _, num := range numbers {
        if num > largest {
            largest = num
        }
        if num < smallest {
            smallest = num
        }
    }

    fmt.Println("Largest number:", largest)
    fmt.Println("Smallest number:", smallest)
}
