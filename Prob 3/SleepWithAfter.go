package main
import ("fmt"
		"time")
 

func main() {
	
	CustomSleep(5)

}
func CustomSleep(duration int){
	fmt.Println("Sleeping for ",duration," seconds")
			c := make(chan string)
			startTime :=time.Now()
				select {
				case  <-c:
					fmt.Println("This case does nothing")
				case <-time.After(time.Duration(duration)*time.Second):
					 elapsedTime :=time.Now().Sub(startTime)
					 fmt.Println("End of sleep .Elapsed time :",elapsedTime,"Press Enter to exit")
				}
		
	var input string
	fmt.Scanln(&input)
}