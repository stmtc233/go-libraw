package metadata

import (
	"bytes"
	"fmt"
)

type LibRawIData struct {
	Make             string
	Model            string
	MakerIndex       uint
	Software         string
	RawCount         uint
	IsFoveon         bool
	DngVersion       uint
	Colors           int
	ColorDescription [5]rune
}

func (idata *LibRawIData) DebugFormat() string {
	var out string

	out += fmt.Sprintf("Make: %s\n", idata.Make)
	out += fmt.Sprintf("Model: %s\n", idata.Model)
	out += fmt.Sprintf("Software: %s\n", idata.Software)

	out += fmt.Sprintf("MakerIndex: %d\n", idata.MakerIndex)
	out += fmt.Sprintf("RawCount: %d\n", idata.RawCount)

	var isFoveon uint = 0
	if idata.IsFoveon {
		isFoveon = 1
	}
	out += fmt.Sprintf("IsFoveon: %d\n", isFoveon)
	out += fmt.Sprintf("DngVersion: %d\n", idata.DngVersion)
	out += fmt.Sprintf("Colors: %d\n", idata.Colors)

	cdesc := string(idata.ColorDescription[:])
	// trim 0s
	if i := bytes.IndexByte([]byte(cdesc), 0); i != -1 {
		cdesc = cdesc[:i]
	}
	out += fmt.Sprintf("Color descriptions: %s\n", cdesc)

	return out
}
