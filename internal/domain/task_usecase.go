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

func (t *Task) GetMetric(title string) (*Metric, bool) {
	for _, m := range t.Metrics {
		if m.Title == title {
			return m, true
		}
	}

	return nil, false
}
