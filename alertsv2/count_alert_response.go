package alertsv2

type AlertCount struct {
	Count int `json:"count"`
}

type CountAlertResponse struct {
	ResponseMeta
	AlertCount AlertCount `json:"data"`
}
