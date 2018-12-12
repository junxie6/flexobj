package flexobj

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func PrintJSON(data interface{}) {
	// produces neatly indented output
	if data, err := json.MarshalIndent(data, "", "    "); err != nil {
		log.Printf("JSON marshaling failed: %s\n", err)
	} else {
		fmt.Printf("%s\n", data)
	}
}

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
