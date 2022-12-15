package main

import "fmt"


func main() {
var number = []int{1,2,3,4,5,6,7,9,10}

for _, evenOdd := range number {
	if evenOdd %2== 0 {
		fmt.Printf( "%d is even \n", evenOdd  )
	}else {
		fmt.Printf("%d is odd \n", evenOdd )


	}
}

}