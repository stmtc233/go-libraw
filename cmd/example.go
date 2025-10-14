package main

import (
	"fmt"
	"image/jpeg"
	"os"

	libraw "github.com/seppedelanghe/go-libraw"
)

const RawPath = "testdata/_SPC2147.NEF"

func main() {
	processor := libraw.NewProcessor(libraw.NewProcessorOptions())
	img, metadata, err := processor.ProcessRaw(RawPath)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}

	err = jpeg.Encode(file, img, &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Camera make: %s\nImage size (h x w): %d x %d\n", metadata.IData.Make, metadata.Sizes.Height, metadata.Sizes.Width)
}
