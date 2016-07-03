
package imageutils

import "fmt"

func NumCells(size uint, stitch_count uint) uint {
	return size * stitch_count
}

func NumPixelsPerCell(totalNumPixels uint, numCells uint) uint {
	if numCells == 0 {
		fmt.Errorf("Cannot compute NumPixelsPerCell with 0 cells.")
		return 0
	}
	return totalNumPixels / numCells
}