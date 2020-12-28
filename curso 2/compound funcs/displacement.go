package main

import "fmt"

func main() {
	var accel float64
	var innitVel float64
	var initialDisp float64
	fmt.Println("Enter acceleration")
	fmt.Scan(&accel)
	fmt.Println("Enter initial velocity")
	fmt.Scan(&innitVel)
	fmt.Println("Enter initial displacement")
	fmt.Scan(&initialDisp)

	displaceFn := GenDisplaceFn(accel, innitVel, initialDisp)

	var time float64
	fmt.Println("Enter time")
	fmt.Scan(&time)
	fmt.Printf("The displacement is: %.4f", displaceFn(time))
}

//GenDisplaceFn calculate displacement
func GenDisplaceFn(acceleration, initialVelocity, initialDispl float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return acceleration*(time*time) + initialVelocity + initialDispl
	}
	return fn
}
