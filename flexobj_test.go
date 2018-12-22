package flexobj_test

import (
	//"fmt"
	"testing"
)
import (
	"github.com/junxie6/flexobj"
)

func TestSetNestedObject(t *testing.T) {
	// Simulating database records
	sqlRow := []struct {
		ExamID       uint32
		ExamName     string
		QuestionID   uint32
		QuestionName string
		ChoiceID     uint32
		ChoiceName   string
		IsSelected   uint8
	}{
		{
			ExamID:       1,
			ExamName:     "this is a 2018 exam",
			QuestionID:   1,
			QuestionName: "What is the best programming language?",
			ChoiceID:     1,
			ChoiceName:   "Go",
			IsSelected:   1,
		},
		{
			ExamID:       1,
			ExamName:     "this is a 2018 exam",
			QuestionID:   1,
			QuestionName: "What is the best programming language?",
			ChoiceID:     2,
			ChoiceName:   "PHP",
			IsSelected:   0,
		},
		{
			ExamID:       1,
			ExamName:     "this is a 2018 exam",
			QuestionID:   2,
			QuestionName: "What is the best editor?",
			ChoiceID:     3,
			ChoiceName:   "Vim",
			IsSelected:   1,
		},
		{
			ExamID:       1,
			ExamName:     "this is a 2018 exam",
			QuestionID:   2,
			QuestionName: "What is the best editor?",
			ChoiceID:     4,
			ChoiceName:   "Visual Studio Code",
			IsSelected:   0,
		},
	}

	//
	flexobj.IsDebug = true

	data := flexobj.New()
	examID := ""
	questionID := ""
	choiceID := ""

	for _, row := range sqlRow {
		examID = flexobj.Uint32ToStr(row.ExamID)
		questionID = flexobj.Uint32ToStr(row.QuestionID)
		choiceID = flexobj.Uint32ToStr(row.ChoiceID)

		if data.IsSet(examID) != true {
			// Initialize an exam object
			exam := flexobj.New()
			exam.Set("ExamID", row.ExamID)
			exam.Set("ExamName", row.ExamName)
			exam.SetArr("QuestionArr", flexobj.New())

			data.SetObj(examID, exam)
		}

		if questionArr := data.GetObj(examID).GetArr("QuestionArr"); questionArr.IsSet(questionID) != true {
			// Initialize a question object
			question := flexobj.New()
			question.Set("QuestionID", row.QuestionID)
			question.Set("QuestionName", row.QuestionName)
			question.SetArr("ChoiceArr", flexobj.New())

			questionArr.SetObj(questionID, question)
		}

		if choiceArr := data.GetObj(examID).GetArr("QuestionArr").GetObj(questionID).GetArr("ChoiceArr"); choiceArr.IsSet(choiceID) != true {
			// Initialize a choice object
			choice := flexobj.New()
			choice.Set("ChoiceID", row.ChoiceID)
			choice.Set("ChoiceName", row.ChoiceName)
			choiceArr.SetObj(choiceID, choice)
		}
	}

	// Print  data in JSON format
	flexobj.PrintJSON(data)
}
