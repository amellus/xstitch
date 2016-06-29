
//This is the main package.
package main

import (
	// Pic "github.com/golang/tour/pic"
	"image"
	// "image/color"
	// "image/draw"
	"os"
	"image/png"
	"flag"
	"fmt"
	// "io"
	// "io/ioutil"
)

var(
	open_filepath string
	write_filepath string
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
	flag.StringVar(&open_filepath, "i", "", "path to image file to be opened")	
	flag.StringVar(&write_filepath, "o", "", "path to new image file to write")
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

func main() {
	f, err := os.Open(open_filepath)

	if err != nil{
		panic(err)
	}

	img, decode_err := png.Decode(f)
	if decode_err != nil {
		panic(decode_err)
	}

	// new_img := image.NewNRGBA(img.Bounds())

	// for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
	// 	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
	// 		fmt.Println(img.At(x,y))
	// 	}
	// }

	// //
	writeImage(img)
}