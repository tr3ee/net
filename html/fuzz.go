package html

import (
	"bytes"
)

func Fuzz(data []byte) int {
	_, err := Parse(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	return 1
}
