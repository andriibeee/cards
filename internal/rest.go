package internal

import (
	"encoding/json"
	"net/http"
)

type ValidatorRequest struct {
	CardNumber string `json:"card_number"`
	ExpMonth   string `json:"exp_month"`
	ExpYear    string `json:"exp_year"`
}

type ValidatorResponse struct {
	Valid  bool     `json:"valid"`
	Causes []string `json:"causes,omitempty"`
}

func MethodHandler(w http.ResponseWriter, r *http.Request) {
	var body ValidatorRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if body.CardNumber == "" || body.ExpMonth == "" || body.ExpYear == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	output := Handle(ValidateCardInput{
		Number: body.CardNumber,
		Month:  body.ExpMonth,
		Year:   body.ExpYear,
	})
	json.NewEncoder(w).Encode(ValidatorResponse{
		Valid:  output.Valid,
		Causes: output.Causes,
	})
}
