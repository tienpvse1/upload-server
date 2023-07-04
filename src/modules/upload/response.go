package upload

import "time"

type UploadInfo struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
}
type UploadResult struct {
	Success bool         `json:"success"`
	Result  []UploadInfo `json:"result"`
}
type UploadResponse struct {
	Success   bool         `json:"success"`
	Status    int          `json:"status"`
	Result    UploadResult `json:"result"`
	Errors    interface{}  `json:"errors"`
	Timestamp time.Time    `json:"timestamp"`
}
