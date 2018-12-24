package flexobj_test

import (
	"fmt"
	"testing"
)
import (
	"github.com/junxie6/flexobj"
)

func TestIteration(t *testing.T) {
	// Exam
	exam := flexobj.New()
	exam.Set("ExamID", 123456)
	exam.Set("ExamName", "this is a 2018 exam")

	// Interation
	for ; exam.Next(); exam.Increase() {
		fmt.Printf("%v: %v\n", exam.Key(), exam.Value())
	}
}
