package solid

import "github.com/go-gl/mathgl/mgl64"

var VertexIntervals = [8]mgl64.Vec3{
	{0, 0, 0},
	{1, 0, 0},
	{1, 1, 0},
	{0, 1, 0},
	{0, 0, 1},
	{1, 0, 1},
	{1, 1, 1},
	{0, 1, 1},
}

var FaceOffsets = [6][4]int{
	{0, 3, 2, 1},
	{4, 5, 6, 7},
	{0, 4, 7, 3},
	{1, 2, 6, 5},
	{0, 1, 5, 4},
	{3, 7, 6, 2},
}
