package model

type (
	Redirect struct {
		ID       int64  `json:"id"`
		From_URL string `json:"from"`
		To_URL   string `json:"to"`
	}
)
