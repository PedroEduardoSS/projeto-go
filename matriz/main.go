package main
import ("fmt")

func main(){
	var matrix = [3][3]int{{3,0,2},{1,2,5},{0,2,4}}
	for index, value := range matrix{
		fmt.Printf("%v\t%v\n", index, value)
	}
	
}