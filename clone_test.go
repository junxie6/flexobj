package flexobj_test

import (
	"fmt"
	"testing"
)
import (
	"github.com/junxie6/flexobj"
)

func TestClone(t *testing.T) {
	// Exam
	exam := flexobj.New()
	exam.Set("ExamID", 123456)
	exam.Set("ExamName", "this is a 2018 exam")

	// Clone data
	examCloned := flexobj.Clone(exam)

	// Set some value
	exam.Set("ExamName", "this is a 2019 exam")

	// Print exam data in JSON format
	fmt.Printf("exam: %s\n", exam.JSONPretty())
	fmt.Printf("examCloned: %s\n", examCloned.JSONPretty())
}
