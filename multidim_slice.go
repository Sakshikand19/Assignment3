package main

import "fmt"
func main(){
	s1:=[] []int {
		{1,2},
		{3,4},
		{5,6},
		{7,8},

	}
	fmt.Println("slice 1:",s1)
	s2:= [] [] string{
		[]string{"str1", "str2"},
		[]string{"str3", "str4"},
		[]string{"str5", "str6"},
	}
		fmt.Println("slice 2:",s2)
		fmt.Println("multidimentional slice s2:")
		for i:=0; i<len(s2); i++{
			fmt.Println(s2[i])
		}
			fmt.Println("slice s2 like matrix:")
			n:= len(s2)
			m:=len(s2[0])
			for i:=0; i<n; i++{
				for j:=0; j<m; j++{
					fmt.Print(s2[i][j]+" ")
					}
					fmt.Print("\n")
				}
			}
		

	

