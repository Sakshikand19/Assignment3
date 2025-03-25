//wap in go language to create and print multidimensional slice
package main

import "fmt"

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    fmt.Println("2D Slice (Matrix):")
    for i := 0; i < len(matrix); i++ {
        fmt.Println(matrix[i])
    }
}
