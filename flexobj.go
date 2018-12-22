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
	Invalid DataType = iota
	Primitive
	HashMap
	OrderedMap
)

var (
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
func (this *FlexObj) Set(key string, val interface{}) {
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

	this.set(key, val, Primitive)
}

// SetObj implies that the value being set is a hash map (object).
func (this *FlexObj) SetObj(key string, val *FlexObj) {
	this.set(key, val, HashMap)
}

// SetArr implies that the value being set is an ordered map (associative array).
func (this *FlexObj) SetArr(key string, val *FlexObj) {
	this.set(key, val, OrderedMap)
}

// set ...
func (this *FlexObj) set(key string, val interface{}, dt DataType) {
	// ReadWrite lock
	this.Lock()
	defer this.Unlock()

	// NOTE: We do not use IsSet() func in order to avoid the extra lock.
	if _, ok := this.data[key]; !ok {
		this.fieldInfoArr = append(this.fieldInfoArr, fieldInfo{
			dataType: dt,
			key:      key,
		})
		this.numOfField++
	}

	this.data[key] = val
}

func (this FlexObj) IsSet(key string) bool {
	// Read lock
	this.RLock()
	defer this.RUnlock()

	//
	_, ok := this.data[key]
	return ok
}

// Get ...
func (this FlexObj) Get(key string) interface{} {
	return this.get(key)
}

// GetObj ...
func (this FlexObj) GetObj(key string) *FlexObj {
	return this.get(key).(*FlexObj)
}

// GetArr ...
func (this FlexObj) GetArr(key string) *FlexObj {
	return this.get(key).(*FlexObj)
}

// get ...
func (this FlexObj) get(key string) interface{} {
	// Read lock
	this.RLock()
	defer this.RUnlock()

	//
	var v interface{}
	var ok bool

	if v, ok = this.data[key]; !ok {
		return nil
	}

	return v
}

func (this FlexObj) JSON() string {
	var err error
	var byteArr []byte

	if byteArr, err = json.Marshal(this); err != nil {
		panic(err.Error())
	}

	return string(byteArr)
}

func (this FlexObj) JSONPretty() string {
	var err error
	var byteArr []byte

	if byteArr, err = json.MarshalIndent(this, "", "    "); err != nil {
		panic(err.Error())
	}

	return string(byteArr)
}
