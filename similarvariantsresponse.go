package model

type SimilarVariantsResponse struct {
	Success         bool  `json:"success"`
	SimilarVariants []int `json:"result"`
}
