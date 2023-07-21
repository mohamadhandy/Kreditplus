package helper

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ImageURL struct {
	FotoKTP    string `json:"foto_ktp"`
	FotoSelfie string `json:"foto_selfie"`
}
