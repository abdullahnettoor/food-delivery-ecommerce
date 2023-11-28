package res

type CommonRes struct {
	Status  string `json:"status,omitempty"`
	Error   any `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Result  any    `json:"result,omitempty"`
}
