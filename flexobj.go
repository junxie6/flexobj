package flexobj

import (
	//"fmt"
	"encoding/json"
	"reflect"
	"sync"
)

// DataType ...
type DataType uint8

const (
	// Invalid ...
	Invalid DataType = iota
	// Primitive ...
	Primitive
	// HashMap ...
	HashMap
	// OrderedMap ...
	OrderedMap
)

var (
	// IsDebug ...
	IsDebug bool
)

// fieldInfo ...
type fieldInfo struct {
	dataType DataType
	key      string
}

// FlexObj ...
type FlexObj struct {
	indexPtr     uint16
	numOfField   uint16
	fieldInfoArr []fieldInfo
	sync.RWMutex // embedded. see http://golang.org/ref/spec#Struct_types
	data         map[string]interface{}
}

// Config ...
// TODO: not yet
type Config struct {
}

// New ...
func New(configArr ...Config) *FlexObj {
	var config Config

	if len(configArr) > 0 {
		config = configArr[0]
	}

	// TODO: not yet
	if config == config {
	}

	// TODO: maybe we can allow user to set up init data when calling New()

	return &FlexObj{
		fieldInfoArr: make([]fieldInfo, 0),
		data:         make(map[string]interface{}),
	}
}

// Set ...
func (fo *FlexObj) Set(key string, val interface{}) {
	if IsDebug == true {
		// NOTE: Add a DataType check only allow Primitive type and panic
		switch reflect.ValueOf(val).Kind() {
		//case reflect.Invalid:
		case reflect.Bool:
		case reflect.Int:
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
		case reflect.Uint:
		case reflect.Uint8:
		case reflect.Uint16:
		case reflect.Uint32:
		case reflect.Uint64:
		//case reflect.Uintptr:
		case reflect.Float32:
		case reflect.Float64:
		//case reflect.Complex64:
		//case reflect.Complex128:
		//case reflect.Array:
		//case reflect.Chan:
		//case reflect.Func:
		//case reflect.Interface:
		// TODO: consider add Map as a Primitive type, but convert Map to *FlexObj (consider reflect cost)
		//case reflect.Map:
		//case reflect.Ptr:
		//case reflect.Slice:
		case reflect.String:
		//case reflect.Struct:
		//case reflect.UnsafePointer:
		default:
			panic("val is not a supported Primitive type")
		}
	}

	fo.set(key, val, Primitive)
}

// SetObj implies that the value being set is a hash map (object).
func (fo *FlexObj) SetObj(key string, val *FlexObj) {
	fo.set(key, val, HashMap)
}

// SetArr implies that the value being set is an ordered map (associative array).
func (fo *FlexObj) SetArr(key string, val *FlexObj) {
	fo.set(key, val, OrderedMap)
}

// set ...
func (fo *FlexObj) set(key string, val interface{}, dt DataType) {
	// ReadWrite lock
	fo.Lock()
	defer fo.Unlock()

	// NOTE: We do not use IsSet() func in order to avoid the extra lock.
	if _, ok := fo.data[key]; !ok {
		fo.fieldInfoArr = append(fo.fieldInfoArr, fieldInfo{
			dataType: dt,
			key:      key,
		})
		fo.numOfField++
	}

	fo.data[key] = val
}

func (fo *FlexObj) IsSet(key string) bool {
	// Read lock
	fo.RLock()
	defer fo.RUnlock()

	//
	_, ok := fo.data[key]
	return ok
}

// Get ...
func (fo *FlexObj) Get(key string) interface{} {
	return fo.get(key)
}

// GetObj ...
func (fo *FlexObj) GetObj(key string) *FlexObj {
	return fo.get(key).(*FlexObj)
}

// GetArr ...
func (fo *FlexObj) GetArr(key string) *FlexObj {
	return fo.get(key).(*FlexObj)
}

// get ...
func (fo *FlexObj) get(key string) interface{} {
	// Read lock
	fo.RLock()
	defer fo.RUnlock()

	//
	var v interface{}
	var ok bool

	if v, ok = fo.data[key]; !ok {
		return nil
	}

	return v
}

// JSON ...
func (fo *FlexObj) JSON() string {
	var err error
	var byteArr []byte

	if byteArr, err = json.Marshal(fo); err != nil {
		panic(err.Error())
	}

	return string(byteArr)
}

// JSONPretty ...
func (fo *FlexObj) JSONPretty() string {
	var err error
	var byteArr []byte

	if byteArr, err = json.MarshalIndent(fo, "", "    "); err != nil {
		panic(err.Error())
	}

	return string(byteArr)
}
