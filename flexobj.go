package flexobj

import (
	//"fmt"
	"reflect"
	"sync"
)

// DataType ...
type DataType uint8

//type HM map[string]interface{}

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

	// TODO: maybe we can allow use to set up init data when calling New()

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
		case reflect.Float32:
		case reflect.Float64:
		case reflect.Array:
		case reflect.Slice:
		case reflect.String:
		case reflect.Struct:
		// TODO: consider add Map as a Primitive type, but convert Map to *FlexObj (consider reflect cost)
		//case reflect.Map:
		default:
			panic("val is not a allowed Primitive type")
		}
	}

	this.set(key, val, Primitive)
}

// SetHM ...
func (this *FlexObj) SetHM(key string, val *FlexObj) {
	this.set(key, val, HashMap)

	// TODO: should we provide "SetM" only? (No more SetHM?)

	// TODO: should we return???
	//return val.(*FlexObj)
}

// SetOM ...
func (this *FlexObj) SetOM(key string, val *FlexObj) {
	this.set(key, val, OrderedMap)

	// TODO: should we return???
	//return val.(*FlexObj)
}

// set ...
func (this *FlexObj) set(key string, val interface{}, dt DataType) {
	// ReadWrite lock
	this.Lock()
	defer this.Unlock()

	// Note: We do not use IsSet() func in order to avoid the extra lock.
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

	_, ok := this.data[key]
	return ok
}

// Get ...
func (this FlexObj) Get(key string) interface{} {
	return this.get(key)
}

// GetHM ...
func (this FlexObj) GetHM(key string) *FlexObj {
	return this.get(key).(*FlexObj)
}

// GetOM ...
func (this FlexObj) GetOM(key string) *FlexObj {
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
		// TODO: we should create if not exist???
		//switch dt {
		//case HashMap, OrderedMap:
		//	return New()
		//default:
		//	return nil
		//}

		return nil
	}

	return v
}
