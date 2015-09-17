package main
import ( "testing"	
		"fmt"
		"time")

//test Sleep with 5 seconds
func TestCustomSleep(t *testing.T){
	startTime :=time.Now()
	CustomSleep(5)
	 var in int
	fmt.Scanln("%d",&in)
	elapsedTime :=time.Now().Sub(startTime)
	fmt.Println("Time elapsed is ",elapsedTime)
	
	if elapsedTime<5.0 {
	   t.Error("Expected 5 seconds,got ",elapsedTime)
	}
	
}