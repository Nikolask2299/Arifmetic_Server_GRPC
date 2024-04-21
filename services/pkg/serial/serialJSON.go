package serial

import (
	"encoding/json"
	"io"
	"net/http"
)

func SerialTaskJSON(request *http.Request) ([]string, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var tasks []string
	if err := json.Unmarshal(body, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
