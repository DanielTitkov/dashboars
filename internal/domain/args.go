package domain

type (
	CreateTaskArgs struct {
		ID          int                    `json:"id"`
		Type        string                 `json:"type"`
		Code        string                 `json:"code"`
		Title       string                 `json:"title"`
		Description string                 `json:"description"`
		Active      bool                   `json:"active"`
		Display     bool                   `json:"display"`
		Schedule    string                 `json:"schedule"`
		Category    string                 `json:"category"`
		Tags        []string               `json:"tags"`
		Args        map[string]interface{} `json:"args"`
	}
)
