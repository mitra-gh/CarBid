package dto

type Otp struct {
	Value string `json:"value"`
	Used  bool   `json:"used"`
}
