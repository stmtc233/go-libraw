package main

import (
	"fmt"
	"os"

	libraw "github.com/stmtc233/go-libraw"
)

const RawPath = `../../_MG_3851.CR2`

func main() {
	// 创建处理器
	processor := libraw.NewProcessor(libraw.NewProcessorOptions())

	fmt.Printf("Extracting thumbnail from %s...\n", RawPath)
	thumb, err := processor.ExtractThumbnail(RawPath)
	if err != nil {
		fmt.Printf("Error extracting thumbnail: %v\n", err)
		return
	}

	fmt.Printf("Thumbnail extracted!\n")
	fmt.Printf("Format: %d (1=JPEG, 2=Bitmap)\n", thumb.Format)
	fmt.Printf("Data size: %d bytes\n", len(thumb.Data))

	if thumb.Format == libraw.ThumbJpeg {
		err = os.WriteFile("thumbnail.jpg", thumb.Data, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println("Saved to thumbnail.jpg")
	} else if thumb.Format == libraw.ThumbBitmap {
		fmt.Println("Thumbnail is in Bitmap format, saving raw data...")
		err = os.WriteFile("thumbnail.raw", thumb.Data, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Unknown thumbnail format")
	}
}
