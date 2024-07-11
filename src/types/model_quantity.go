package types

type Quantity struct {
	Count int32 `json:"count,omitempty"`

	Measure Scalar `json:"measure,omitempty"`
}
