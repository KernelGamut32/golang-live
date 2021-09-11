package main

// START OMIT
import (
	"fmt"  // we import packages in order to use them
	"math" // we'll use the math package too
)

func main() {
	var name, age = "The Hamburglar", 34
	fmt.Printf("Name is %s\n", name)
	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Printf("The square root of %d is %f\n", age, math.Sqrt(float64(age)))
}

// END OMIT
