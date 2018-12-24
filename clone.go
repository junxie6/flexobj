package flexobj

// Clone ...
func Clone(src *FlexObj) *FlexObj {
	dstNew := New()
	clone(dstNew, src)
	return dstNew
}

// clone ...
func clone(dst *FlexObj, src *FlexObj) {
	src.Reset()

	for ; src.Next(); src.Increase() {
		switch src.DataType() {
		case Primitive:
			dst.Set(src.Key(), src.Value())
		case HashMap:
			dstNew := New()
			dst.SetObj(src.Key(), dstNew)
			clone(dstNew, src.Value().(*FlexObj))
		case OrderedMap:
			dstNew := New()
			dst.SetArr(src.Key(), dstNew)
			clone(dstNew, src.Value().(*FlexObj))
		default:
			panic("Unsupported data type")
		}
	}

	src.Reset()
}
