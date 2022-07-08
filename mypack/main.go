package main

import (
	"fmt"
	"mypack/mymodulo"
)

func main(){
	run := true
	var status int = 0
	
	for run {
		mymodulo.PrintHello()
		fmt.Print("Digite um valor:")
		fmt.Scanf("%v", &status)
		if status == 1{
			break
		}
	}
}