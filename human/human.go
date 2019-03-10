package human

import (
	"errors"
	"math/rand"
)

var (
	// TOOHIGHERROR will be the error message when guessed too high
	TOOHIGHERROR = "too high"
	// TOOLOWERROR will be the error message when guessed too low
	TOOLOWERROR = "too low"
)

// Person will be the type of person
type Person struct {
	hiddenAge int
}

// NewPerson will initialize person
func NewPerson() *Person {
	age := rand.Intn(101)

	person := Person{}
	person.hiddenAge = age
	return &person
}

// GuessAge will guess the age of the person
func (p *Person) GuessAge(age int) error {
	if age > p.hiddenAge {
		return errors.New(TOOHIGHERROR)
	}

	return nil
}
