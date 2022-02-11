package domain

import (
	"encoding/json"
	"fmt"
)

func (t *Task) JSONString() string {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}`, err)
	}
	return string(jsonBytes)
}
