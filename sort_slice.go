package main
import (
	"fmt"
	"sort"
)
func main(){
	   intSlice :=[]int{42, 23, 16, 15,8,4}
	   fmt.Println("Before:",intSlice)

	   sort.Ints(intSlice)
	   fmt.Println("After:",intSlice)
}
