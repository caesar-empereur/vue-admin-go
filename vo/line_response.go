package vo

type LineResponse struct {
	IndexCount int64 `json:"indexCount"`

	SellCount int64 `json:"sellCount"`

	SecCount int64 `json:"secCount"`

	CurrTime int64 `json:"currTime"`
}
