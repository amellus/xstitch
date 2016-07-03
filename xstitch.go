
//This is the main package.
package main

import (
	"image"
	// "image/color"
	// "image/draw"
	"os"
	"image/png"
	"flag"
	"fmt"
	"github.com/amellus/xstitch/imageutils"
)

var(
	open_filepath string
	write_filepath string
	stitch_count uint
	width uint
	height uint
)

//Makes some pixel values
func Pixels(dx, dy int) [][]uint8 {
	
	pic := make([][]uint8, dy)
	for y:= 0; y < dy; y++ {
		pic[y] = make([]uint8, dx)
		for x:=0; x < dx; x++ {
			pic[y][x] = uint8(x^y)
		}
	}
	return pic
}

func init() {
	//TODO: input validation
	flag.StringVar(&open_filepath, "i", "", "path to image file to be opened")	
	flag.StringVar(&write_filepath, "o", "ximage.png", "path to new image file to write")
	flag.UintVar(&stitch_count, "s", 14, "desired stitch count")
	flag.UintVar(&width, "w", 10, "desired image width")
	flag.UintVar(&height, "h", 15, "desired image height")
	flag.Parse()
}

func writeImage(img image.Image) {
	f, err := os.Create(write_filepath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Writing to", write_filepath, "...")
	defer f.Close()
	png.Encode(f, img)
	
}

var(
	num_cells_horizontal uint
	num_cells_vertical uint
	num_pixels_per_cell_horizontal uint
	num_pixels_per_cell_vertical uint
)

func main() {
	f, err := os.Open(open_filepath)

	if err != nil{
		panic(err)
	}

	img, decode_err := png.Decode(f)
	if decode_err != nil {
		panic(decode_err)
	}

	initDimensions(img.Bounds())

	avg := imageutils.AveragePixelValue(400, 700, int(num_pixels_per_cell_horizontal), int(num_pixels_per_cell_vertical), img)

	fmt.Printf("%v\n", avg)

	// xmin, xmax, ymin, ymax := rect.Min.X, rect.Max.X, rect.Min.Y, rect.Max.Y 
	// for y := ymax; y >= ymin; y-- {
	// 	for x := xmin; x < xmax; x++ {
	// 		fmt.Printf("x=%v, y=%v, color=%s\n", x, y, img.At(x,y))
	// 	}
	// }

	// new_img := image.NewNRGBA(img.Bounds())

	// for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
	// 	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
	// 		fmt.Println(img.At(x,y))
	// 	}
	// }

	// //
	writeImage(img)
}

func initDimensions(rect image.Rectangle) {
	xmin, xmax, ymin, ymax := uint(rect.Min.X), uint(rect.Max.X), uint(rect.Min.Y), uint(rect.Max.Y) 
	num_pixels_horizontal := xmax - xmin
	num_pixels_vertical := ymax - ymin

	fmt.Printf("num_pixels_horizontal = %v, num_pixels_vertical = %v.\n", num_pixels_horizontal, num_pixels_vertical)

	num_cells_horizontal = imageutils.NumCells(width, stitch_count)
	num_cells_vertical = imageutils.NumCells(height, stitch_count)

	num_pixels_per_cell_horizontal = imageutils.NumPixelsPerCell(num_pixels_horizontal, num_cells_horizontal)
	num_pixels_per_cell_vertical = imageutils.NumPixelsPerCell(num_pixels_vertical, num_cells_vertical)

	fmt.Printf("num_cells_horizontal = %v, num_cells_vertical = %v, num_pixels_per_cell_horizontal = %v, num_pixels_per_cell_vertical = %v.\n",
		num_cells_horizontal, num_cells_vertical, num_pixels_per_cell_horizontal, num_pixels_per_cell_vertical)
}