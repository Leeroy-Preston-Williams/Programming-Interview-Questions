// @Author:		Leeroy P. Williams
// @Date:		20/05/19
// @Desc:		Should display Fizz for multiples of3 & Buzz for multiples of 5
//				then display FizzBuzz for multiples of both

package main

import "fmt"

func main(){
	
	for i := 100; i > 0; i--{
		if (i%3==0) && (i%5==0){
			fmt.Println("FizzBuzz")
		}else if i%3==0{
			fmt.Println("Fizz")
		}else if i%5==0{
			fmt.Println("Buzz")
		}else{
		fmt.Println(i)
	}
	}
}