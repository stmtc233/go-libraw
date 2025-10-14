package metadata

import "time"

type ImgMetadata struct {
	CaptureTimestamp int64
	CaptureDate      time.Time

	IData LibRawIData
	Sizes LibRawSizes
}
