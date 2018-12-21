package flexobj

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"reflect"
)

// MarshalJSON implements Marshaler
// NOTE: https://golang.org/pkg/encoding/json/#Marshaler
func (this FlexObj) MarshalJSON() ([]byte, error) {
	//
	var byteArr []byte
	var err error

	buf := bytes.NewBufferString("{")
	length := this.numOfField - 1

	count := uint16(0)
	var fi fieldInfo

	for _, fi = range this.fieldInfoArr {
		//fmt.Printf("dataType: %v %v\n", fi.key, fi.dataType)

		switch fi.dataType {
		case Primitive:
			if byteArr, err = json.Marshal(this.Get(fi.key)); err != nil {
				return nil, err
			}
		case HashMap:
			if byteArr, err = json.Marshal(this.GetObj(fi.key)); err != nil {
				return nil, err
			}
		case OrderedMap:
			if byteArr, err = json.Marshal(this.GetArr(fi.key).Values()); err != nil {
				return nil, err
			}
		default:
			return nil, err
		}

		buf.WriteString(fmt.Sprintf(`"%s":%s`, fi.key, string(byteArr)))

		if count < length {
			buf.WriteString(",")
		}

		count++
	}

	buf.WriteString("}")

	return buf.Bytes(), nil
}
