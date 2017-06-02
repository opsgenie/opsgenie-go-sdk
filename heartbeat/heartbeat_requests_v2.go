package heartbeat

type AddHeartbeatRequestV2 struct {
	Name         string `json:"name"`
	Interval     int    `json:"interval"`
	IntervalUnit string `json:"intervalUnit"`
	Enabled      bool  `json:"enabled"`
	Description  string `json:"description,omitempty"`
}

type UpdateHeartbeatRequestV2 struct {
	Interval     int    `json:"interval,omitempty"`
	IntervalUnit string `json:"intervalUnit,omitempty"`
	Enabled      bool  `json:"enabled,omitempty"`
	Description  string `json:"description,omitempty"`
}


type PingHeartbeatRequest struct {
	Name   string `json:"name"`
}