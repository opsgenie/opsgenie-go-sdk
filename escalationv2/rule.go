package escalationv2

type Rule struct {
	Condition  string    `json:"condition"`
	NotifyType string    `json:"notifyType"`
	Delay      Delay     `json:"delay"`
	Recipient  Recipient `json:"recipient"`
}

type Delay struct {
	TimeAmount int    `json:"timeAmount"`
	TimeUnit   string `json:"timeUnit"`
}

type Recipient struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
