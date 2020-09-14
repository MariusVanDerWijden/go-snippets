package snap

import (
	"fmt"

	"github.com/golang/snappy"
)

var maxUint24 = ^uint32(0) >> 8

func Fuzz(data []byte) int {
	size, err := snappy.DecodedLen(data)
	if err != nil {
		return 1
	}
	if size > int(maxUint24) {
		return -1
	}
	out, err := snappy.Decode(nil, data)
	if err != nil {
		return 1
	}
	if size != len(out) {
		panic(fmt.Sprintf("Size: %v %v", len(out), size))
	}
	return 1
}
