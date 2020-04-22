package upload

type UploadResponse struct {
	Url      string `json:"url"`
	Path     string `json:"path"`
	FileName string `json:"fileName"`
}
