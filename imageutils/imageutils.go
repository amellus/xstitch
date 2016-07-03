
package imageutils

import (
	"fmt"
	"math"
	"image"
	"image/color"
)

func NumCells(size uint, stitch_count uint) uint {
	return size * stitch_count
}

func NumPixelsPerCell(totalNumPixels uint, numCells uint) uint {
	if numCells == 0 {
		fmt.Errorf("Cannot compute NumPixelsPerCell with 0 cells.\n")
		return 0
	}
	return totalNumPixels / numCells
}

func AveragePixelValue(cell_x int, cell_y int, num_pixels_x int, num_pixels_y int, img image.Image) color.NRGBA {
	total_pixels := uint32(num_pixels_x*num_pixels_y)
	if total_pixels == 0 {
		fmt.Errorf("Cannot get average pixel value for cell with no pixels.\n")
		// zero := uint8(0)
		return color.NRGBA{0,0,0,0}
	}

	r, g, b, a := uint32(0), uint32(0), uint32(0), uint32(0)
	y_start, x_start := img.Bounds().Min.Y + cell_y, img.Bounds().Min.X + cell_x
	x_end := int(math.Min(float64(x_start + num_pixels_x), float64(img.Bounds().Max.X-1)))
	y_end := int(math.Min(float64(y_start + num_pixels_y), float64(img.Bounds().Max.Y-1)))
	for y := y_start; y < y_end; y++ {
		for x := x_start; x < x_end; x++ {
			pixel_r, pixel_g, pixel_b, pixel_a := img.At(x,y).RGBA()
			r += pixel_r
			g += pixel_g
			b += pixel_b
			a += pixel_a
		}
	}
	
	r = r / total_pixels
	g = g / total_pixels
	b = b / total_pixels
	a = a / total_pixels
	return color.NRGBA {uint8(r),uint8(g),uint8(b),uint8(a)}
}