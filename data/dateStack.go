package data

import (
	"fmt"
	"time"
)

// DateNode is the specific date of contact.
type DateNode struct {
	Date time.Time
	Last *DateNode
}

// Stack contains the top DateNode for the stack.
type Stack struct {
	Top  *DateNode
	Size int
}

// Push allows the user to push a new DateNode onto the Stack.
func (s *Stack) Push(date time.Time) {
	newDate := &DateNode{
		Date: date,
		Last: nil,
	}

	if s.Top == nil {
		s.Top = newDate
	} else {
		newDate.Last = s.Top
		s.Top = newDate
	}
	s.Size++
}

func (s *Stack) printStack() {
	if s.Top == nil {
		fmt.Println("No dates found")
	}

	currentDate := s.Top
	for {
		if s.Size == 1 {
			fmt.Printf("%v - Date of Last Contact\n", currentDate.Date.Format("02 Jan 2006"))
			break
		} else if currentDate == s.Top {
			fmt.Printf("%v - Date of Last Contact\n", currentDate.Date.Format("02 Jan 2006"))
		} else if currentDate.Last == nil {
			fmt.Println(currentDate.Date.Format("02 Jan 2006"))
			break
		} else {
			fmt.Println(currentDate.Date.Format("02 Jan 2006"))
		}
		currentDate = currentDate.Last
	}
}
