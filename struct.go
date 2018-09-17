package main

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmhMultiples float64 = 1.60934

type car struct {
	gasPedal      uint16 // min 0 max 255
	brakePedal    uint16
	steeringWheel int16 // -32k - +32K
	topSpeedKmh   float64
}

// value receiver method is used for computation
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

// value receiver method is used for computation
func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmhMultiples)
}

// pointer receiver method is used for computation
func (c *car) updateSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

func main() {
	dreamCar := car{
		gasPedal:      65000,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKmh:   225.00}

	fmt.Println(dreamCar.gasPedal)
	fmt.Println(dreamCar.kmh())
	fmt.Println(dreamCar.mph())

	// update speed via pointer receiver method
	dreamCar.updateSpeed(500)
	fmt.Println(dreamCar.kmh())
	fmt.Println(dreamCar.mph())
}
