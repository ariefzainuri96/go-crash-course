package l7pointerreceivers

import (
	"fmt"
)

type car struct {
	color string
}

// Pointer Receivers
func (c *car) setColor(color string) {
	c.color = color
}

// Non-Pointer Receivers
// func (c car) setColor(color string) {
// 	c.color = color
// }

/*
We often used pointer receivers to modify the value of a variable.
*/

func suchMain() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
}
