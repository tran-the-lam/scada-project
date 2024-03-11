package service

type LoginInfo struct {
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	DeviceID  string `json:"device_id"`
	Time      string `json:"time"`
}

type User struct {
	UserID   string `json:"user_id"`
	Role     string `json:"role"`
	Password string `json:"password,omitempty"`
	Status   int    `json:"status"`
}

type Event struct {
	EventName   string  `json:"event"`
	ParameterID string  `json:"parameter_id"` // convert sensor_id to parameter_id
	Parameter   string  `json:"parameter"`
	Value       float64 `json:"value"`
	Threshold   float64 `json:"threshold"`
	Timestamp   uint64  `json:"timestamp"`
}

type ChangePwd struct {
	UpdatedAt int64 `json:"updated_at"`
}
