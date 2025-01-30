package types

type ResultResponse struct {
	ValueA float64 `json:"valueA"`
	ValueB float64 `json:"valueB"`
	Result float64 `json:"result"`
}

func NewResultResponse() *ResultResponse {

	return &ResultResponse{}
}
func (r *ResultResponse) SetValues(valueA, valueB float64) *ResultResponse {

	r.ValueA = valueA
	r.ValueB = valueB
	return r
}

func (r *ResultResponse) SetResult(value float64) *ResultResponse {

	r.Result = value
	return r
}
