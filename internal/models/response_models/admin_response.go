package res

type AdminLoginRes struct {
	Status  string `json:"status,omitempty"`
	Token   string `json:"token,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
