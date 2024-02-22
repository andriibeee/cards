package internal

import "time"

type ValidateCardInput struct {
	Number string
	Month  string
	Year   string
}

type ValidateCardOutput struct {
	Valid  bool
	Causes []string
}

func Handle(input ValidateCardInput) ValidateCardOutput {
	causes := make([]string, 0)

	validNum := CheckCreditCardNumber(input.Number)
	if !validNum {
		causes = append(causes, "Invalid credit card number")
	}

	validMonth := CheckCreditCardMonth(input.Month)
	if !validMonth {
		causes = append(causes, "Invalid expiration month")
	}

	expired := HasCardExpired(input.Month, input.Year, time.Now())
	if expired {
		causes = append(causes, "Credit card has expired")
	}

	return ValidateCardOutput{
		Valid:  validNum && validMonth && !expired,
		Causes: causes,
	}
}
