package pkg

import "encoding/json"

type BandwidthStat struct {
	Interface     string  `json:"interface"`
	UploadSpeed   float64 `json:"upload"`
	DownloadSpeed float64 `json:"download"`
}

func (n BandwidthStat) String() string {
	s, _ := json.Marshal(n)
	return string(s)
}
