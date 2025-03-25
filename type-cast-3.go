package main
import "fmt"
func main(){
	var x float64=3.14159
	var y int=int(x)//convert float64 to int
	fmt.Println("x(float64)",x)
	fmt.Println("y (int):",y)

	var a int = 42
	var b float64 = float64(a)

	
	fmt.Println("a (int):",a)
	fmt.Println("b (float64):",b)

}