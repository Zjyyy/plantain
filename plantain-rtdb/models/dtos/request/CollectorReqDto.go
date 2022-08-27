package request

type CollectorReq struct {
	CollectorName string `json:"collectorName"`
	Version       string `json:"version"`
	DllPath       string `json:"dllPath"`
	ConnStr       string `json:"connStr"`
	Setting       string `json:"setting"`
	Des           string `json:"des"`
	RtTableName   string `json:"rtTableName"`
}
