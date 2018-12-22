package flexobj

import (
	"fmt"
	"log"
	"reflect"
)

func Decode(src *FlexObj, dst interface{}) error {
	return decode(reflect.ValueOf(src.data), reflect.ValueOf(dst).Elem())
}

func decode(src reflect.Value, dst reflect.Value) (err error) {
	switch src.Kind() {
	case reflect.Ptr:
		// TODO: there may be a more efficient way (instead of convert it back to interface)
		if err = decode(reflect.ValueOf(src.Interface().(*FlexObj).data), dst); err != nil {
			return err
		}
	case reflect.Struct:
		// TODO: to be implemented
		fmt.Printf("src.Kind() is reflect.Struct - to be implemented\n")
	case reflect.Slice:
		dst.Set(reflect.MakeSlice(dst.Type(), src.Len(), src.Cap()))

		for i := 0; i < src.Len(); i++ {
			if err = decode(src.Index(i).Elem(), dst.Index(i)); err != nil {
				return err
			}
		}
	case reflect.Map:
		for _, key := range src.MapKeys() {
			srcVal := src.MapIndex(key)
			dstVal := dst.FieldByName(key.String())

			if dstVal.IsValid() != true {
				// TODO: add some message
				log.Printf("Error: dstVal is invalid\n")
				continue
			}

			if dstVal.CanSet() != true {
				// TODO: add some message

				continue
			}

			//fmt.Printf("======== Field: %s ===\n", key.String())

			if srcVal.Elem().Type() != dstVal.Type() {
				// TODO: to be implemented

				// Compare a custom type
				switch srcVal.Elem().Type() {
				case reflect.TypeOf((*FlexObj)(nil)):
					// NOTE: Values() return one result.
					questionArr := srcVal.Elem().MethodByName("Values").Call([]reflect.Value{})[0]

					// TODO: not fully understand. Find out how to handle []reflect.Value. It seems unnecessary (Use questionArr instead)
					//questionVal := reflect.ValueOf(questionArr.Interface().([]interface{}))

					if err = decode(questionArr, dstVal); err != nil {
						return err
					}

					//fmt.Printf("dstVal: %v: %v: %#v\n", dstVal.Kind(), dstVal.Type(), dstVal)

					//tryStruct := reflect.ValueOf(dstVal)
					//tryStruct := reflect.Indirect(reflect.ValueOf(dstVal))
					//tryStruct := reflect.Indirect(dstVal)
					//tryStruct := dstVal.Type()
					//sliceType := tryStruct.Elem()

					//yyy := tryStruct.Type()
					//fmt.Printf("yyy: %v\n", yyy)

					//typeOfT := tryStruct.Type()

					//fmt.Printf("tryStruct: %v: %v: %#v\n", tryStruct.Kind(), tryStruct.Type(), tryStruct)
					//fmt.Printf("tryStruct: %v:  %#v\n", tryStruct.Kind(), tryStruct.NumField())
					//fmt.Printf("sliceType: %v:  %#v\n", sliceType.Kind(), sliceType.NumField())

					//for i := 0; i < sliceType.NumField(); i += 1 {
					//	//	//	//valueField := tryStruct.Field(i)
					//	typeField := sliceType.Field(i)
					//	fmt.Printf("FN: %v\n", typeField.Name)
					//}

					//reflect.MakeSlice(dstVal.Type(), dstVal.Len(), dstVal.Cap())
				default:
					return fmt.Errorf("Type mismatched: %v vs %v\n", srcVal.Elem().Type(), dstVal.Type())
				}

				continue
			}

			//dstTyp, _ := dst.Type().FieldByName(key.String())
			//fmt.Printf("T %v\n", dstTyp.Name)

			//fmt.Printf("srcVal: %v: %v: %#v\n", srcVal.Elem().Kind(), srcVal.Elem().Type(), srcVal.Elem())
			//fmt.Printf("dstVal: %v: %v: %#v\n", dstVal.Kind(), dstVal.Type(), dstVal)

			if err = decode(srcVal.Elem(), dstVal); err != nil {
				return err
			}

			//switch srcVal.Elem().Kind() {
			//case reflect.Uint32:
			//	fmt.Printf("It is Uint32\n")
			//case reflect.String:
			//	fmt.Printf("It is String\n")
			//}
		}
	case reflect.String:
		// TODO: Review
		// NOTE: https://gist.github.com/hvoecking/10772475
		//translatedString := dict[src.Interface().(string)]
		//dst.SetString(translatedString)

		//fmt.Printf("It is String\n")
		dst.Set(src)
	default:
		dst.Set(src)
	}

	return nil
}
