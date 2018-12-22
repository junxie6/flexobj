# FlexObj

FlexObj helps storing the nested objects as easy as PHP associative array and maintaining the insertion order

## Features:
* Maintain insertion order
* Output in JSON format
* Decode to a struct
* Deep clone a FlexObj
* Iteration

## Examples:

### Storing the nested objects and output in JSON format

main.go:

```
package main

import (
	"github.com/junxie6/flexobj"
)

func main() {
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
}
```

JSON output:

```
{
    "ExamID": 1,
    "ExamName": "this is a 2018 exam",
    "QuestionArr": [
        {
            "QuestionID": 1,
            "QuestionName": "What is the best programming language?",
            "ChoiceArr": [
                {
                    "ChoiceID": 1,
                    "ChoiceName": "Go"
                },
                {
                    "ChoiceID": 2,
                    "ChoiceName": "PHP"
                }
            ]
        }
    ]
}
```
