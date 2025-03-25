// Go program to illustrate
// the use of booleans
package main
import "fmt"

func main() {
	
	// variables
str1 := "sakshikand"
str2:= "SAKSHIKAND"
str3:= "sakshikand"
result1:= str1 == str2
result2:= str1 == str3
	
// Display the result
fmt.Println( result1)
fmt.Println( result2)
	
// Display the type of 
// result1 and result2
fmt.Printf("The type of result1 is %T and "+
				"the type of result2 is %T",
							result1, result2)
	
}
