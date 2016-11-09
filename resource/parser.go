package resource

import (
	"encoding/json"
	"fmt"
	"io"
)

func ParseInput(inputStream io.Reader) (Input, error) {
	var input Input
	if err := json.NewDecoder(inputStream).Decode(&input); err != nil {
		return input, fmt.Errorf("parsing input: %s", err.Error())
	}
	return input, nil
}
