package internal

import (
	"reflect"
	"testing"
)

func mustEqual(tb testing.TB, got, want interface{}) {
	tb.Helper()

	if !reflect.DeepEqual(got, want) {
		tb.Fatalf("got: %v, want: %v", got, want)
	}
}

func TestHandler(t *testing.T) {
	t.Run("Valid cards", func(t *testing.T) {
		expectedOutput := ValidateCardOutput{
			Valid:  true,
			Causes: []string{},
		}
		mustEqual(t, Handle(ValidateCardInput{
			Number: "4111111111111111",
			Month:  "12",
			Year:   "28",
		}), expectedOutput)
		mustEqual(t, Handle(ValidateCardInput{
			Number: "4111111111111111",
			Month:  "12",
			Year:   "2028",
		}), expectedOutput)
	})

	t.Run("Invalid cards", func(t *testing.T) {
		mustEqual(t, Handle(ValidateCardInput{
			Number: "4111111111111111",
			Month:  "01",
			Year:   "2021",
		}), ValidateCardOutput{
			Valid:  false,
			Causes: []string{"Credit card has expired"},
		})
		mustEqual(t, Handle(ValidateCardInput{
			Number: "1111111111111",
			Month:  "10",
			Year:   "2021",
		}), ValidateCardOutput{
			Valid:  false,
			Causes: []string{"Invalid credit card number", "Credit card has expired"},
		})
		mustEqual(t, Handle(ValidateCardInput{
			Number: "4111111111111111",
			Month:  "1",
			Year:   "21",
		}), ValidateCardOutput{
			Valid:  false,
			Causes: []string{"Credit card has expired"},
		})
		mustEqual(t, Handle(ValidateCardInput{
			Number: "1111111111111",
			Month:  "1",
			Year:   "28",
		}), ValidateCardOutput{
			Valid:  false,
			Causes: []string{"Invalid credit card number"},
		})
	})
}
