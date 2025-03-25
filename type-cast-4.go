package main
import "fmt"
func main(){
	var floatNum float64 =3.14159
	var intNum int =int(floatNum)
	fmt.Println("float value ",floatNum)
	fmt.Println("int value(after casting)",intNum)
}

