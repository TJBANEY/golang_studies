package codeWars

// CartesianNeighbor - A Cartesian coordinate system is a coordinate system that specifies each point uniquely in a plane by a pair of numerical coordinates,
// which are the signed distances to the point from two fixed perpendicular directed lines, measured in the same unit of length.
// The сoordinates of a point in the grid are written as (x,y). Each point in a coordinate system has eight neighboring points. Provided that the grid step = 1.
// It is necessary to write a function that takes a coordinate on the x-axis and y-axis and returns a list of all the neighboring points. Points inside list
// dont have to been sorted (any order is valid).
func CartesianNeighbor(x, y int) [][]int {
	// Slice that holds int slices. This is what will be returned
	neighbors := [][]int{}

	// Array of fixed length 8 that holds slices of type int
	relativeCoords := [8][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

	for _, relCoord := range relativeCoords {
		xAxis := x + relCoord[0]
		yAxis := y + relCoord[1]

		neighbor := []int{xAxis, yAxis}
		neighbors = append(neighbors, neighbor)
	}

	return neighbors

	// Example Result
	// CartesianNeighbor(2,2) -> {{1,1},{1,2},{1,3},{2,1},{2,3},{3,1},{3,2},{3,3}}
}
