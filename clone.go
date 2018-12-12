package flexobj

import (
	"fmt"
)

func Clone(src *FlexObj) *FlexObj {
	dstNew := New()
	clone(dstNew, src)
	return dstNew
}

func clone(dst *FlexObj, src *FlexObj) {
	src.Reset()

	for ; src.Next(); src.Increase() {
		fmt.Printf("%v (%v): %v\n", src.Key(), src.DataType(), src.Value())

		switch src.DataType() {
		case Primitive:
			dst.Set(src.Key(), src.Value())
		case HashMap:
			dstNew := New()
			dst.SetHM(src.Key(), dstNew)
			clone(dstNew, src.Value().(*FlexObj))
		case OrderedMap:
			dstNew := New()
			dst.SetOM(src.Key(), dstNew)
			clone(dstNew, src.Value().(*FlexObj))
		default:
			panic("Unsupported data type")
		}
	}

	src.Reset()
}
