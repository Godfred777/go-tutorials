package main

import (
	"fmt"

	"example.com/tutorials/tutorial-1/area"

	"example.com/tutorials/tutorial-1/perimeter"
)

func main() {
	fmt.Println("Welcome to this console application")

	var option int
	fmt.Print("What do you want to calculate. \n1. Area Of A Sector. \n2. Area Of A Rectangle. \n3. Area Of A Trapezium. \n4. Area Of A Triangle. \n5. Arc length of an Arc. \n6. Perimeter Of A Quadrilatual.\n")
	fmt.Scan(&option)
	switch option {
	case 1:
		var (
			rad float64
			ang float64
		)
		fmt.Print("Enter the values for the radius and the angle: ")
		fmt.Scan(&rad, &ang)
		res := area.AreaOfSector(rad, ang)
		fmt.Print(res)
	case 2:
		var (
			len float64
			bre float64
		)
		fmt.Print("Enter the values for the length and breadth: ")
		fmt.Scan(&len, &bre)
		res := area.AreaOfRectangle(len, bre)
		fmt.Print(res)
	case 3:
		var (
			h float64
			s float64
			l float64
		)
		fmt.Print("Enter the values for the height, short length and long length: ")
		fmt.Scan(&h, &s, &l)
		res := area.AreaOfTrapezium(h, s, l)
		fmt.Print(res)
	case 4:
		var (
			b float64
			h float64
		)
		fmt.Print("Enter the values for the base and height: ")
		fmt.Scan(&b, &h)
		res := area.AreaOfTriangle(b, h)
		fmt.Print(res)
	case 5:
		var (
			r float64
			a float64
		)
		fmt.Print("Enter the values for the radius and the angle: ")
		fmt.Scan(&r, &a)
		res := perimeter.ArcLength(r, a)
		fmt.Print(res)
	case 6:
		var (
			fi float64
			sc float64
			th float64
			fo float64
		)
		fmt.Print("Enter the values for the lengths: ")
		fmt.Scan(&fi, &sc, &th, &fo)
		res := perimeter.Perimeter(fi, sc, th, fo)
		fmt.Print(res)
	default:
		fmt.Println("This number is invalid")
	}
}
