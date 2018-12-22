package flexobj

import (
	"strconv"
)

// StrToUint32 converts uint32 string to uint32 integer
func StrToUint32(num string) uint32 {
	var n uint64
	var err error

	if n, err = strconv.ParseUint(num, 10, 32); err != nil {
		return 0
	}

	return uint32(n)
}

// Uint32ToStr converts uint32 integer to string
func Uint32ToStr(num uint32) string {
	return strconv.FormatUint(uint64(num), 10)
}
