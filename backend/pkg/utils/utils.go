package utils

import (
	"encoding/json"

	"github.com/google/uuid"
)

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func GenerateUUID() uuid.UUID {
	return uuid.New()
}
