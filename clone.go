package flexobj

import (
	"fmt"
	"reflect"
	//"unsafe"
)

func (this FlexObj) GetIndexPtr() uint16 {
	return this.indexPtr
}
func (this FlexObj) GetNumOfField() uint16 {
	return this.numOfField
}
func (this *FlexObj) SetNumOfField(n uint16) {
	this.numOfField = n
}

func Clone(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	src := reflect.ValueOf(v)
	dst := reflect.New(src.Type()).Elem()

	//return clone(dst, src).(*FlexObj)

	clone(dst, src)

	return dst.Interface()
}

func clone(dst reflect.Value, src reflect.Value) {
	switch src.Kind() {
	case reflect.Ptr:
		srcVal := src.Elem()

		if !src.IsValid() {
			return
		}

		dst.Set(reflect.New(srcVal.Type()))

		clone(dst.Elem(), srcVal)
	case reflect.Slice:
		fmt.Printf("not yet\n")
	case reflect.Struct:
		srcType := src.Type()

		(dst.Interface().(FlexObj)).SetNumOfField(uint16(1))
		fff := dst.Interface().(FlexObj).GetNumOfField()
		fmt.Printf("fff: %v\n", fff)

		//xxx := srcVal.Elem().MethodByName("Values").Call([]reflect.Value{})
		//xxx := src.MethodByName("GetIndexPtr").Call([]reflect.Value{})
		//xxx := src.MethodByName("GetNumOfField").Call([]reflect.Value{})[0].Uint()

		//fmt.Printf("HERE: %#v\n", xxx)
		//for _, questionArr := range xxx {
		//	bbb := questionArr.Uint()
		//	fmt.Printf("HERE: %v\n", bbb)
		//}

		//dstMethod := dst.MethodByName("SetNumOfField")

		//params := make([]reflect.Value, dstMethod.Type().NumIn())
		//params[0] = src.MethodByName("GetNumOfField").Call([]reflect.Value{})[0]
		//dstMethod.Call(params)

		for i := 0; i < src.NumField(); i++ {
			//if srcType.Field(i).PkgPath != "" {
			//	continue
			//}

			switch srcType.Field(i).Name {
			case "numOfField":
			}

			//fmt.Printf("FN: %v %v\n", asdf.Name, src.Field(i).CanSet())

			//fmt.Printf("FN: %v\n", src.Field(i).Kind())
			//clone(dst.Field(i), src.Field(i))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		//asdf := src.Uint()
		//fmt.Printf("ASDF: %#v\n", asdf)
		//dst.SetUint(src.Uint())
		//dst.SetUint(5)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		//dst.SetInt(src.Int())
	case reflect.String:
	default:
		//asdf := unsafe.Pointer(src.UnsafeAddr())

		//fmt.Printf("OK: %#v\n", (asdf))

		//asdf := src

		//asdf := reflect.New(src.Type())
		//dst.Set(asdf)
	}
}
