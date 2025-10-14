package metadata

import "fmt"

type LibRawSizes struct {
	RawHeight uint16
	RawWidth  uint16
	Height    uint16
	Width     uint16
	Iheight   uint16
	Iwidth    uint16
}

func (sizes *LibRawSizes) DebugFormat() string {
	var out string

	out += fmt.Sprintf("RawHeight: %d\n", sizes.RawHeight)
	out += fmt.Sprintf("RawWidth: %d\n", sizes.RawWidth)
	out += fmt.Sprintf("Height: %d\n", sizes.Height)
	out += fmt.Sprintf("Width: %d\n", sizes.Width)
	out += fmt.Sprintf("IHeight: %d\n", sizes.Iheight)
	out += fmt.Sprintf("IWidth: %d\n", sizes.Iwidth)

	return out
}
