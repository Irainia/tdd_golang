package human

import (
	"testing"
)

func TestGuessAgeTooHigh(t *testing.T) {
	person := NewPerson()
	actual := person.GuessAge(person.hiddenAge + 1)
	expeced := TOOHIGHERROR

	if actual == nil {
		t.Errorf("expected: %s - actual: nil", expeced)
		return
	}

	if actual.Error() != expeced {
		t.Errorf("expected: %s - actual: %s", expeced, actual.Error())
	}
}

func TestGuessAgeTooLow(t *testing.T) {
	person := NewPerson()
	actual := person.GuessAge(person.hiddenAge - 1)
	expeced := TOOLOWERROR

	if actual == nil {
		t.Errorf("expected: %s - actual: nil", expeced)
		return
	}

	if actual.Error() != expeced {
		t.Errorf("expected: %s - actual: %s", expeced, actual.Error())
	}
}

func TestGuessAgeCorrect(t *testing.T) {
	person := NewPerson()
	actual := person.GuessAge(person.hiddenAge)

	if actual != nil {
		t.Errorf("expected: nil - actual: %s", actual.Error())
		return
	}
}
