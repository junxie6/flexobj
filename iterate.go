package flexobj

import (
//"fmt"
)

// Key ...
func (this *FlexObj) Key() string {
	return this.fieldInfoArr[this.indexPtr].key
}

// DataType ...
func (this *FlexObj) DataType() DataType {
	return this.fieldInfoArr[this.indexPtr].dataType
}

// Value implements iterator
func (this *FlexObj) Value() interface{} {
	// Read lock
	this.RLock()
	defer this.RUnlock()

	//
	return this.data[this.Key()]
}

// Next implements iterator
func (this *FlexObj) Next() bool {
	if this.indexPtr < this.numOfField {
		return true
	}

	return false
}

func (this *FlexObj) Increase() {
	this.indexPtr++
}

// Err implements iterator
func (this *FlexObj) Err() error {
	// TODO: to be implemented
	return nil
}

// Values ...
func (this *FlexObj) Values() []interface{} {
	this.Reset()

	valArr := make([]interface{}, 0, this.numOfField)

	for ; this.Next(); this.Increase() {
		valArr = append(valArr, this.Value())
	}

	this.Reset()
	return valArr
}

// Reset sets the internal pointer of an array to its first element
func (this *FlexObj) Reset() {
	this.indexPtr = 0
}
