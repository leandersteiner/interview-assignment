package calculator

type AdditionRequest struct {
	SummandOne float64 `json:"summand_one"`
	SummandTwo float64 `json:"summand_two"`
}

type AdditionResponse struct {
	Sum float64 `json:"sum"`
}

type SubtractionRequest struct {
	Minuend    float64 `json:"minuend"`
	Subtrahend float64 `json:"subtrahend"`
}

type SubtractionResponse struct {
	Difference float64 `json:"difference"`
}

type MultiplicationRequest struct {
	FactorOne float64 `json:"factor_one"`
	FactorTwo float64 `json:"factor_two"`
}

type MultiplicationResponse struct {
	Product float64 `json:"product"`
}

type DivisionRequest struct {
	Dividend float64 `json:"dividend"`
	Divisor  float64 `json:"divisor"`
}

type DivisionResponse struct {
	Quotient float64 `json:"quotient"`
}

type RecentResponse struct {
	Results  []string `json:"calculations"`
	Metadata Metadata `json:"pagination"`
}
