package golibraw

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stmtc233/go-libraw/pkg/metadata"
)

const testPath = "./testdata"
const cppVersion = "./test_metadata"

func getAllFilesInTestDir() []string {
	entries, err := os.ReadDir(testPath)
	if err != nil {
		panic(err)
	}
	paths := make([]string, len(entries))
	for i, e := range entries {
		paths[i] = filepath.Join(testPath, e.Name())
	}
	return paths
}

func compareToC(path string, meta metadata.ImgMetadata) bool {
	cmd := exec.Command(cppVersion, path)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running C++ program needed to compare LibRaw outputs:", err)
		return false
	}

	outStr := string(output)
	idataDebug := meta.IData.DebugFormat()
	sizesDebug := meta.Sizes.DebugFormat()
	var combined string = idataDebug + sizesDebug

	eq := outStr == combined
	if !eq {
		fmt.Printf("C output:\n%s\n", outStr)
		fmt.Printf("Go output:\n%s", combined)
	}

	return eq

}

// TestProcessRaw uses ProcessRaw to decode the RAW file into an image.Image and checks the metadata.
func TestProcessRaw(t *testing.T) {
	processor := NewProcessor(NewProcessorOptions())

	for _, path := range getAllFilesInTestDir() {
		img, meta, err := processor.ProcessRaw(path)
		if err != nil {
			t.Fatalf("ProcessRaw failed: %v", err)
		}
		if img == nil {
			t.Fatal("ProcessRaw returned a nil image")
		}

		bounds := img.Bounds()
		if bounds.Dx() <= 0 || bounds.Dy() <= 0 {
			t.Errorf("Invalid image dimensions: %v", bounds)
		}

		if meta.CaptureTimestamp == 0 || meta.CaptureDate.IsZero() {
			t.Error("ProcessRaw returned invalid metadata")
		}

		if !compareToC(path, meta) {
			t.Errorf("Metadata returned from C != Go for '%s'", path)
		}
	}
}

// TestConcurrentProcessRaw runs ProcessRaw concurrently in multiple goroutines.
func TestConcurrentProcessRaw(t *testing.T) {
	processor := NewProcessor(NewProcessorOptions())

	paths := getAllFilesInTestDir()
	var wg sync.WaitGroup
	wg.Add(len(paths))

	for i, path := range getAllFilesInTestDir() {
		go func(idx int) {
			defer wg.Done()
			img, meta, err := processor.ProcessRaw(path)
			if err != nil {
				t.Errorf("Goroutine %d: ProcessRaw failed: %v", idx, err)
				return
			}
			if img == nil {
				t.Errorf("Goroutine %d: returned nil image", idx)
			}
			if meta.CaptureTimestamp == 0 || meta.CaptureDate.IsZero() {
				t.Errorf("Goroutine %d: returned invalid metadata", idx)
			}

			if !compareToC(path, meta) {
				t.Errorf("Metadata returned from C != Go for '%s'", path)
			}

		}(i)
	}
	wg.Wait()
}
