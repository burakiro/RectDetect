package main

import (
	"fmt"
	gocv "gocv.io/x/gocv"
	"image"
	"image/color"
)

func main() {
	window := gocv.NewWindow("Rectangle Detection")
	defer window.Close()

	input := gocv.NewMat()
	defer input.Close()
	filename := "Rectangle.png"

	input = gocv.IMRead(filename, gocv.IMReadColor)



	gray := gocv.NewMat()
	defer gray.Close()

	gocv.CvtColor(input, &gray, gocv.ColorBGRToGray)

	mask := gocv.NewMat()
	defer mask.Close()

	gocv.Threshold(gray, &mask, 0, 255, gocv.ThresholdBinaryInv|gocv.ThresholdOtsu)
	contours := gocv.FindContours(mask, gocv.RetrievalExternal, gocv.ChainApproxSimple)

	fmt.Println(mask.Size())
	var biggestContourIdx int = -1
	var biggestContourArea float64 = 0

	drawing := gocv.NewMat()
	defer drawing.Close()

	drawing = gocv.NewMatWithSize(mask.Size()[0], mask.Size()[1], gocv.MatTypeCV8UC3)
	for i, _ := range contours {
		color := color.RGBA{0, 255, 0, 0}
		gocv.DrawContours(&drawing, contours, i, color, -1)
		var ctArea float64 = gocv.ContourArea(contours[i])
		if ctArea > biggestContourArea {
			biggestContourArea = ctArea
			biggestContourIdx = i
		}

	}
	if biggestContourIdx < 0 {
		fmt.Println("No Contour Found!")

	}
	var boundingBox gocv.RotatedRect
	boundingBox = gocv.MinAreaRect(contours[biggestContourIdx])

	corners := make([]image.Point, 4)


	boundingBox.Contour = corners
	fmt.Println(boundingBox.Contour)

	//gocv.Line(&drawing, corners[0], corners[1], color.RGBA{255, 255, 255, 0}, -1)
	//gocv.Line(&drawing, corners[1], corners[2], color.RGBA{255, 255, 255, 0}, -1)
	//gocv.Line(&drawing, corners[2], corners[3], color.RGBA{255, 255, 255, 0}, -1)
	//gocv.Line(&drawing, corners[3], corners[0], color.RGBA{255, 255, 255, 0}, -1)



	window.IMShow(input)
	gocv.WaitKey(0)
	window.IMShow(drawing)
	gocv.WaitKey(0)

	gocv.WaitKey(0)
}
