package flexobj

import (
//"fmt"
)

// Key ...
func (fo *FlexObj) Key() string {
	return fo.fieldInfoArr[fo.indexPtr].key
}

// DataType ...
func (fo *FlexObj) DataType() DataType {
	return fo.fieldInfoArr[fo.indexPtr].dataType
}

// Value implements iterator
func (fo *FlexObj) Value() interface{} {
	// Read lock
	fo.RLock()
	defer fo.RUnlock()

	//
	return fo.data[fo.Key()]
}

// Next implements iterator
func (fo *FlexObj) Next() bool {
	if fo.indexPtr < fo.numOfField {
		return true
	}

	return false
}

// Increase ...
func (fo *FlexObj) Increase() {
	fo.indexPtr++
}

// Err implements iterator
func (fo *FlexObj) Err() error {
	// TODO: to be implemented
	return nil
}

// Values ...
func (fo *FlexObj) Values() []interface{} {
	fo.Reset()

	valArr := make([]interface{}, 0, fo.numOfField)

	for ; fo.Next(); fo.Increase() {
		valArr = append(valArr, fo.Value())
	}

	fo.Reset()
	return valArr
}

// Reset sets the internal pointer of an array to its first element
func (fo *FlexObj) Reset() {
	fo.indexPtr = 0
}
