//Problem 2
package main

import ("fmt" 
"math")

type Shape interface{
	area() float64
	perimeter() float64
}

type Circle struct{
	x,y,radius float64
}

type Rectangle struct{
	x1,y1,x2,y2 float64
}
func main() {
	rect :=Rectangle{3,4,6,8}
	cir  :=Circle{3,4,10}
	fmt.Println("perimeter of Rectangle is",perim(rect))
	fmt.Println("Area of Rectangle is",area(rect))
	fmt.Println("Area of Circle is",area(cir))
	fmt.Println("perimeter of Circle is",perim(cir))
}

func area(shape Shape) float64{
	return shape.area()
}
func  perim(shape Shape) float64{
	return shape.perimeter()
}

func distance(x1,y1,x2,y2 float64) float64{

	xdif := x2 - x1
	ydif := y2 - y1
	return math.Sqrt(xdif*xdif + ydif*ydif)
}

////Implementing interface for area() of Rectangle
func (rect Rectangle) area() float64{
	len := distance(rect.x1,rect.y1,rect.x1,rect.y2)
	wid := distance(rect.x1,rect.y1,rect.x2,rect.y1)

	return len * wid
} 

//Implementing interface for perimeter() of rectangle 
func (rect Rectangle)perimeter() float64{
	len := distance(rect.x1,rect.y1,rect.x1,rect.y2)
	wid := distance(rect.x1,rect.y1,rect.x2,rect.y1)

	return len *2+ wid*2
} 

//Implementing interface for area() of Circle 
func (cir Circle)area() float64{
	if cir.radius >=0{
	return math.Pi*cir.radius*cir.radius
}else{
		fmt.Println("Negative radius.Exiting..")
		return -1
	} 
} 

//Implementing interface for perimeter() of Circle 
func (cir Circle)perimeter() float64{
	if cir.radius >=0{
	return 2*math.Pi*cir.radius
	}else{
		fmt.Println("Negative radius.Exiting..")
		return -1
	} 
} 



