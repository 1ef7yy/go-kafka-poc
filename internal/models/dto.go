package models

type RedisRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ValidPhoneResponse struct {
	Status     bool   `json:"status"`
	Normalized string `json:"normalized"`
}

type InvalidPhoneResponse struct {
	Status bool `json:"status"`
}
