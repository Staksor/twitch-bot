package structs

type GptRequest struct {
	Application string `json:"application"`
	Instance    string `json:"instance"`
	Message     string `json:"message"`
}
