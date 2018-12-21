package flexobj_test

import (
	//"fmt"
	"testing"
)
import (
	"github.com/junxie6/flexobj"
)

type Exam struct {
	ExamID      uint32
	ExamName    string
	QuestionArr []struct {
		QuestionID   uint32
		QuestionName string
		ChoiceArr    []struct {
			ChoiceID   uint32
			ChoiceName string
		}
	}
}

func TestDecode(t *testing.T) {
	//
	exam_ExamID := "1"
	exam := flexobj.New()

	exam.Set("ExamID", flexobj.StrToUint32(exam_ExamID))
	exam.Set("ExamName", "this is a 2018 exam")

	// QuestionArr
	questionArr := flexobj.New()
	exam.SetArr("QuestionArr", questionArr)

	// Question 1
	q1_QuestionID := "1"
	q1 := flexobj.New()
	questionArr.SetObj(q1_QuestionID, q1)

	q1.Set("QuestionID", flexobj.StrToUint32(q1_QuestionID))
	q1.Set("QuestionName", "What is the best programming language?")

	// Question 1 - ChoiceArr
	q1_ChoiceArr := flexobj.New()
	q1.SetArr("ChoiceArr", q1_ChoiceArr)

	// Question 1 - Choice 1
	q1c1_ChoiceID := "1"
	q1c1 := flexobj.New()
	q1_ChoiceArr.SetObj(q1c1_ChoiceID, q1c1)

	q1c1.Set("ChoiceID", flexobj.StrToUint32(q1c1_ChoiceID))
	q1c1.Set("ChoiceName", "Go")

	// Question 1 - Choice 2
	q1c2_ChoiceID := "2"
	q1c2 := flexobj.New()
	q1_ChoiceArr.SetObj(q1c2_ChoiceID, q1c2)

	q1c2.Set("ChoiceID", flexobj.StrToUint32(q1c2_ChoiceID))
	q1c2.Set("ChoiceName", "PHP")

	// Print exam data in JSON format
	flexobj.PrintJSON(exam)

	// Decode to exam2 struct
	exam2 := Exam{}
	flexobj.Decode(exam, &exam2)

	// Print exam2 data in JSON format
	flexobj.PrintJSON(exam2)
}
