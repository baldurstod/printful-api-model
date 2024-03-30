package model

type ProductResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Product `json:"pf"`
	} `json:"result"`
}
