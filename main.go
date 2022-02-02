package main

import "fmt"

type Employee struct {
	FirstName, LastName string
	Salary              Salary
}

type Salary float64

func (s Salary) AfterRaise() float64 {
	return CountRaise(float64(s))
}

func main() {
	e := Employee{
		FirstName: "Mike",
		LastName:  "Hello",
		Salary:    Salary(2000),
	}
	fmt.Printf("Employee: %s %s\nBasic Salary:\t%.2f USD\n", e.FirstName, e.LastName, e.Salary)
	fmt.Printf("After Raise:\t%.2f USD\n", e.Salary.AfterRaise())
}

func CountRaise(salary float64) float64 {
	var raise float64

	switch {
	case salary >= 2000:
		raise = (float64(salary) * 15.5) / 100
	case salary < 2000:
		raise = float64(salary*11) / 100
	}

	return salary + raise
}
