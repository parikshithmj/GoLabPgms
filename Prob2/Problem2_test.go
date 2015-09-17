package main
import ("testing"
		"math")

func TestCirclePerimeter(t *testing.T) {
	cir  :=Circle{3,4,10}
	output := perim(cir)
	expectedVal := math.Pi*2*10
	if output!=expectedVal{
	   t.Error("Expected :",expectedVal," got",output)
	}

}

func TestCirclePerimeterNegativeRadius(t *testing.T) {
	cir  :=Circle{-3,4,-10}
	output := perim(cir)
	expectedVal := -1.0
	if output!=expectedVal{
	   t.Error("Expected :",expectedVal," got",output)
	}

}
func TestCircleArea(t *testing.T) {
	cir  :=Circle{3,4,10}
	output := area(cir)
	expectedVal := math.Pi*10*10
	if output!=expectedVal{
	   t.Error("Expected :",expectedVal," got",output)
	}

}

func TestCircleAreaNegativeRadius(t *testing.T) {
	cir  :=Circle{3,4,-10}
	output := area(cir)
	expectedVal := -1.0
	if output!=expectedVal{
	   t.Error("Expected :",expectedVal," got",output)
	}

}

func TestRectanglePerimeter(t *testing.T) {
	rect :=Rectangle{3,4,6,8}
	output := perim(rect)
	expectedVal := float64(3*2+4*2)
	if output!=expectedVal{
	   t.Error("Expected :",expectedVal," got",output)
	}
}
func TestRectanglePerimeterNegativeCordinates(t *testing.T) {
	rect :=Rectangle{-3,4,-6,8}
	output := perim(rect)
	expectedVal := float64(3*2+4*2)
	if output!=expectedVal{
	   t.Error("Expected :",expectedVal," got",output)
	}
}
func TestRectangleArea(t *testing.T) {
	rect :=Rectangle{3,4,6,8}
	output := area(rect)
	expectedVal := float64(3*4)
	if output!=expectedVal{
	   t.Error("Expected :",expectedVal," got",output)
	}
}
func TestRectangleAreaNegativeCordinates(t *testing.T) {
	rect :=Rectangle{-3,4,-6,8}
	output := area(rect)
	expectedVal := float64(3*4)
	if output!=expectedVal{
	   t.Error("Expected :",expectedVal," got",output)
	}
}


