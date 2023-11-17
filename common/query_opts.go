package common

type QueryOpts struct {
	Page    int    `json:"page"`
	Size    int    `json:"limit"`
	OrderBy string `json:"order_by"`
	Order   string `json:"order"`
}
