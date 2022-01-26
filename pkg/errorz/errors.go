package errorz

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Errors struct {
	Name        string `json:"name"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func NewError(name, message, desc string) []byte {
	data, err := json.Marshal(Errors{
		Name:        name,
		Message:     message,
		Description: desc,
	})
	if err != nil {
		log.Fatal("error marshal failure")
		return nil
	}
	return data
}
