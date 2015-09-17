//Problem 1 
package main
import "fmt"

func main(){
	fmt.Println("Enter a positive integer :")
	var number int
	fmt.Scanf("%d",&number)
	fmt.Println("The Fibnocci number is for ",number,"is",GetFibbonaci(number))
	
}
//Only for positive input,based on the expression mentioned in the question is for positive numbers.
//input : number
//output :Fibnocci number
func GetFibbonaci(number int) int {
	if number <0 {
		fmt.Println("Invalid input.Exiting..")
		return -1
	}else if number ==1|| number ==0 {
	    return number
	}else{
		return GetFibbonaci(number - 1)+ GetFibbonaci(number - 2)
	}
}

