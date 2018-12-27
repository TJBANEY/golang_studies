package code_wars

import (
	"fmt"
)

func code_wars(){
	fmt.Println("Test")
}

// A Cartesian coordinate system is a coordinate system that specifies each point uniquely in a plane by a pair of numerical coordinates, 
// which are the signed distances to the point from two fixed perpendicular directed lines, measured in the same unit of length.
// The сoordinates of a point in the grid are written as (x,y). Each point in a coordinate system has eight neighboring points. Provided that the grid step = 1.
// It is necessary to write a function that takes a coordinate on the x-axis and y-axis and returns a list of all the neighboring points. Points inside list
// dont have to been sorted (any order is valid).
func CartesianNeighbor(x, y int) [][]int {
	intArrayOne := []int{}
	intArrayTwo := []int{}
	intArrayArray := [][]int{intArrayOne, intArrayTwo}


	// [[-1, +1], [-1, -1], [+1, +1], [+1, -1], [+1, 0], [-1, 0], [0, -1], [0, +1]]

	// Arrays are immutable in size, perhaps completely?
	// Slices can have items added or removed but are still references to some underlying array.
	return intArrayArray

	// Example Result
	// CartesianNeighbor(2,2) -> {{1,1},{1,2},{1,3},{2,1},{2,3},{3,1},{3,2},{3,3}}
}