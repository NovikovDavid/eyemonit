package pkg

type Bandwidth struct {
	Interface     string  `json:"interface"`
	UploadSpeed   float64 `json:"upload"`
	DownloadSpeed float64 `json:"download"`
}
