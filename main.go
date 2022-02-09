package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

type Employee struct {
	FirstName, LastName string
	Salary              float64
}

func main() {
	rand.Seed(time.Now().Unix())
	e := RandomEmployee()
	fmt.Printf("Employee: %s %s, Salary: %.2f USD\n", e.FirstName, e.LastName, e.Salary)
}

func CountRaise(salary float64) float64 {
	var raise float64

	switch {
	case salary >= 5000:
		raise = (float64(salary) * 11.5) / 100
	case salary < 5000:
		raise = float64(salary*15.5) / 100
	}

	return salary + raise
}

func RandomEmployee() Employee {
	return Employee{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Salary:    RandomSalary(),
	}
}

func RandomSalary() float64 {
	return rand.Float64() * 10000
}
