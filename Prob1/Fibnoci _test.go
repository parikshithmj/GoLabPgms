package main
import "testing"
func TestFibbonaciPositiveInteger(number *testing.T){
	 var output int
	 output = GetFibbonaci(7)
	 if output!=13{
	 	number.Error("Expected 13,but got",output)
	 }
}

func TestFibbonaciZero(number *testing.T){
	 var output int
	 output = GetFibbonaci(0)
	 if output!=0{
	 	number.Error("Expected 0,but got",output)
	 }
}

func TestFibbonaciNegativeNo(number *testing.T){
	 var output int
	 output = GetFibbonaci(-8)
	 if output!=-1{
	 	number.Error("Expected -1,but got",output)
	 }
}