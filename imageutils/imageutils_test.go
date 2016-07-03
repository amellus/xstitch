
package imageutils

import(
	"testing"
)

type num_cells_testcase struct {
	size uint
	stitch_count uint
	num_cells uint
}

var num_cells_testcases = []num_cells_testcase{
	{10, 14, 140},
	{15, 14, 210},
	{0, 4, 0},
	{9, 0, 0},
}

func TestNumCells(t *testing.T) {
	for _, c := range num_cells_testcases {
		v := NumCells(c.size, c.stitch_count)
		if v != c.num_cells {
			t.Error("For size", c.size,
			 "stitch_count", c.stitch_count,
			 "expected", c.num_cells,
			 "got", v)
		}
	}
}

type pixels_per_cell_testcase struct {
	total_pixels uint
	num_cells uint
	pixels_per_cell uint
}

var pixels_per_cell_testcases = []pixels_per_cell_testcase {
	{1000, 10, 100},
	{0, 9, 0},
	{1000, 6, 166},
	{1, 8, 0},
	{9, 7, 1},
}

func TestNumPixelsPerCell(t *testing.T) {
	for _, c := range pixels_per_cell_testcases {
		v := NumPixelsPerCell(c.total_pixels, c.num_cells)
		if v != c.pixels_per_cell {
			t.Error("For total_pixels", c.total_pixels,
			 "num_cells", c.num_cells,
			 "expected", c.pixels_per_cell,
			 "got", v)
		}
	}
}

// type average_pixel_testcase struct {
// 	Img image.Image
// 	X uint
// 	Y uint
// 	Average color.Color
// }

// var uniform_red = image.Image.NewUniform(image.NRGBA{255, 0, 0, 1})

// var average_pixel_testcases = []average_pixel_testcase {
// 	{uniform_red, 0, 0, uniform_red},
// }

// func TestAveragePixelValue(t *testing.T) {
// 	for _, c := range average_pixel_testcases {
// 		v := AveragePixelValue(c.X, c.Y, c.Img)
// 		if v != c.Average {
// 			t.Error("For x =", c.X, "y =", c.Y)
// 		}
// 	}

// }