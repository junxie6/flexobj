# FlexObj

[![Build Status](https://travis-ci.org/junxie6/flexobj.svg?branch=master)](https://travis-ci.org/junxie6/flexobj)
[![codecov](https://codecov.io/gh/junxie6/flexobj/branch/master/graph/badge.svg)](https://codecov.io/gh/junxie6/flexobj)
[![Go Report Card](https://goreportcard.com/badge/github.com/junxie6/flexobj)](https://goreportcard.com/report/github.com/junxie6/flexobj)
[![GoDoc](https://godoc.org/github.com/junxie6/flexobj?status.svg)](https://godoc.org/github.com/junxie6/flexobj)

FlexObj makes storing the nested objects as easy as PHP associative array and it maintains the insertion order. It is most useful when handling the SQL database result set in golang.

## Why?

I wanted to store the SQL database result set as a list of the nested objects in golang while maintaining the insertion order (especially when generating the reports). Then, outputs the data as a JSON string to the browser (JavaScript). Go's built-in hashmap does not keep the insertion order and require type assertion when retrieving the nested objects (which makes the codes look more verbose). So, I built this tool to resolve my needs.

## Features:
- Maintain the insertion order
- Output a FlexObj instance as a JSON string (MarshalJSON implements Marshaler)
- Decode from a FlexObj instance to a struct (could be used to ensure the correct data type is used)
- Deep clone from a FlexObj instance
- Insertion-order iteration over the FlexObj instance
- No type assertion is required when retrieving a nested object
- Go 1.11 Modules support

## Installation

1. Download and install FlexObj:

```sh
$ go get -u github.com/junxie6/flexobj
```

2. Import FlexObj in your code:

```go
import (
	"github.com/junxie6/flexobj"
)
```

## Examples:

### Set some primitive type data (boolean, integer, float, string)

exampleSetPrimitiveType.go:
```go
package main

import (
	"fmt"
)
import (
	"github.com/junxie6/flexobj"
)

func main() {
	data := flexobj.New()
	data.Set("ExamID", 1)
	data.Set("ExamName", "this is a 2018 exam")
	data.Set("IsDone", false)

	// Print data in JSON format
	fmt.Printf("Output: %s\n", data.JSONPretty())
}
```

JSON output:
```json
{
    "ExamID": 1,
    "ExamName": "this is a 2018 exam",
    "IsDone": false
}
```

### Set a hashmap data

exampleSetHashMap.go:
```go
package main

import (
	"fmt"
)
import (
	"github.com/junxie6/flexobj"
)

func main() {
	data := flexobj.New()

	user := flexobj.New()
	user.Set("UserID", 5)
	user.Set("UserName", "Bot 1")

	data.SetObj("User", user)

	// Print data in JSON format
	fmt.Printf("Output: %s\n", data.JSONPretty())
}
```

JSON output:
```json
{
    "User": {
        "UserID": 5,
        "UserName": "Bot 1"
    }
}
```

### Set a ordered map data

exampleSetOrderedMap.go:
```go
package main

import (
	"fmt"
)
import (
	"github.com/junxie6/flexobj"
)

func main() {
	data := flexobj.New()
	questionArr := flexobj.New()
	data.SetArr("QuestionArr", questionArr)

	q1 := flexobj.New()
	q1.Set("QuestionID", 1)
	q1.Set("QuestionName", "What is the best programming language?")
	questionArr.SetObj("1", q1)

	q2 := flexobj.New()
	q2.Set("QuestionID", 2)
	q2.Set("QuestionName", "What is the best editor?")
	questionArr.SetObj("2", q2)

	// Print data in JSON format
	fmt.Printf("Output: %s\n", data.JSONPretty())
}
```

JSON output:
```json
{
    "QuestionArr": [
        {
            "QuestionID": 1,
            "QuestionName": "What is the best programming language?"
        },
        {
            "QuestionID": 2,
            "QuestionName": "What is the best editor?"
        }
    ]
}
```

### Store the database result set and output it as a JSON string

exampleDatabaseResultSet.go:
```go
package main

import (
	"fmt"
)
import (
	"github.com/junxie6/flexobj"
)

func main() {
	// Simulating a database result set
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
			choice.Set("IsSelected", row.IsSelected)

			choiceArr.SetObj(choiceID, choice)
		}
	}

	// Print data in JSON format
	fmt.Printf("Output: %s\n", data.JSONPretty())
}
```

JSON output:
```json
{
    "1": {
        "ExamID": 1,
        "ExamName": "this is a 2018 exam",
        "QuestionArr": [
            {
                "QuestionID": 1,
                "QuestionName": "What is the best programming language?",
                "ChoiceArr": [
                    {
                        "ChoiceID": 1,
                        "ChoiceName": "Go",
                        "IsSelected": 1
                    },
                    {
                        "ChoiceID": 2,
                        "ChoiceName": "PHP",
                        "IsSelected": 0
                    }
                ]
            },
            {
                "QuestionID": 2,
                "QuestionName": "What is the best editor?",
                "ChoiceArr": [
                    {
                        "ChoiceID": 3,
                        "ChoiceName": "Vim",
                        "IsSelected": 1
                    },
                    {
                        "ChoiceID": 4,
                        "ChoiceName": "Visual Studio Code",
                        "IsSelected": 0
                    }
                ]
            }
        ]
    }
}
```

### Store the nested objects and output it as a JSON string

exampleNestedObject.go:
```go
package main

import (
	"fmt"
)
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
	q1c1.Set("IsSelected", 1)

	// Question 1 - Choice 2
	q1c2_ChoiceID := "2"
	q1c2 := flexobj.New()
	q1_ChoiceArr.SetObj(q1c2_ChoiceID, q1c2)

	q1c2.Set("ChoiceID", flexobj.StrToUint32(q1c2_ChoiceID))
	q1c2.Set("ChoiceName", "PHP")
	q1c2.Set("IsSelected", 0)

	// Print exam data in JSON format
	fmt.Printf("Output: %s\n", exam.JSONPretty())
}
```

JSON output:
```json
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
                    "ChoiceName": "Go",
                    "IsSelected": 1
                },
                {
                    "ChoiceID": 2,
                    "ChoiceName": "PHP",
                    "IsSelected": 0
                }
            ]
        }
    ]
}
```

### Clone

exampleClone.go:
```go
package main

import (
	"fmt"
)
import (
	"github.com/junxie6/flexobj"
)

func main() {
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
```

Output:
```
exam: {
    "ExamID": 123456,
    "ExamName": "this is a 2019 exam"
}
examCloned: {
    "ExamID": 123456,
    "ExamName": "this is a 2018 exam"
}
```

### Iteration

exampleIteration.go:
```go
package main

import (
	"fmt"
)
import (
	"github.com/junxie6/flexobj"
)

func main() {
	// Exam
	exam := flexobj.New()
	exam.Set("ExamID", 123456)
	exam.Set("ExamName", "this is a 2018 exam")

	// Interation
	for ; exam.Next(); exam.Increase() {
		fmt.Printf("%v: %v\n", exam.Key(), exam.Value())
	}
}
```

Output:
```
ExamID: 123456
ExamName: this is a 2018 exam
```
