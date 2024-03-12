package model

type CountriesResponse struct {
	Success   bool      `json:"success"`
	Countries []Country `json:"result"`
}

type Country struct {
	Name   string  `json:"name"`
	Code   string  `json:"code"`
	Region string  `json:"region"`
	States []State `json:"states"`
}

type State struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
