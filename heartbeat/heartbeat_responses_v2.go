package heartbeat

type HeartbeatResponseV2 struct {
	Code      int `json:"code"`
	Data      HeartbeatData `json:"data"`
	Took      float32 `json:"took"`
	RequestId string   `json:"requestId"`
}

type HeartbeatData struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Enabled       bool   `json:"enabled"`
	Interval      int    `json:"interval"`
	IntervalUnit  string `json:"intervalUnit"`
	Expired       bool   `json:"expired"`
}

type HeartbeatMetaResponseV2 struct {
	Code      int `json:"code"`
	Data      HeartbeatMetaData `json:"data"`
	Took      float32 `json:"took"`
	RequestId string   `json:"requestId"`
}

type HeartbeatMetaData struct {
	Name          string `json:"name"`
	Enabled       bool   `json:"enabled"`
	Expired       bool   `json:"expired"`
}

type PingHeartbeatResponse struct {
	Result  	string `json:"result"`
	Took      float32 `json:"took"`
	RequestId string   `json:"requestId"`
}


